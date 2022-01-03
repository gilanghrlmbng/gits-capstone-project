package controllers

import (
	"math/rand"
	"net/http"
	"src/api/models"
	"src/entity"
	"src/utils"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

func CreateOrder(c echo.Context) error {
	var items []entity.ItemOrder
	ord := new(entity.Order)

	if err := c.Bind(&items); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	ord.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)

	if claims.UserId == "" && claims.User != "warga" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	ord.IdWarga = claims.UserId
	ord.CreatedAt = time.Now()

	var hargaTotal int64 = 0
	var idKelProduk string = ""
	for idx, item := range items {
		produk, err := models.GetProdukByID(c, item.IdProduk)
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		if idKelProduk == "" {
			idKelProduk = produk.IdKeluarga
		} else if produk.IdKeluarga != idKelProduk {
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusBadRequest,
				Message: "Item yang dipesan harus dari toko yang sama",
			})
		}

		if err := item.ValidateCreate(); err.Code > 0 {
			return utils.ResponseError(c, err)
		}

		items[idx].HargaTotal = produk.Harga * item.Jumlah
		hargaTotal += items[idx].HargaTotal

		entropys := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
		items[idx].Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropys).String()

		items[idx].IdOrder = ord.Id
		items[idx].CreatedAt = ord.CreatedAt
	}
	c.Logger().Info("Items ", items)
	ord.Harga_total = hargaTotal
	ord.Status = entity.OrderStatusDipesan
	ord.IdKeluargaPenjual = idKelProduk

	// pengecekan jenis pembayaran
	if c.Param("pembayaran") != "Saldo" && c.Param("pembayaran") != "COD" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Jenis pembayaran tidak terdaftar",
		})
	}

	entropy = ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	pembayaran := entity.Pembayaran{
		Id:                ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String(),
		IdOrder:           ord.Id,
		IdKeluargaPembeli: claims.IdKeluarga,
		IdKeluargaPenjual: idKelProduk,
		Jumlah_pembayaran: hargaTotal,
		Jenis:             c.Param("pembayaran"),
		CreatedAt:         ord.CreatedAt,
	}

	if c.Param("pembayaran") == "Saldo" {
		dompet, err := models.GetDompetKeluargaByID(c, "", claims.IdKeluarga)
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		if dompet.Jumlah < hargaTotal {
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusBadRequest,
				Message: entity.SaldoTidakCukup,
			})
		}

		pembayaran.Status = entity.PembayaranTerbayar

		dompet.Jumlah = dompet.Jumlah - hargaTotal
		_, err = models.UpdateDompetKeluargaById(c, dompet.Id, &dompet)
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	} else {
		pembayaran.Status = entity.PembayaranBelumDibayar
	}

	_, err := models.CreateOrder(c, ord)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.CreateBatchItemOrder(c, items)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.CreatePembayaran(c, &pembayaran)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	orders, _ := models.GetOrderByID(c, ord.Id)

	return utils.ResponseDataOrder(c, utils.JSONResponseDataOrder{
		Code:        http.StatusCreated,
		CreateOrder: orders,
		Message:     "Berhasil",
	})
}

func GetAllOrderPembeli(c echo.Context) error {
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)

	var status string
	if c.QueryParam("status") == "0" {
		status = entity.OrderStatusCancel
	} else if c.QueryParam("status") == "1" {
		status = entity.OrderStatusDipesan
	} else if c.QueryParam("status") == "2" {
		status = entity.OrderStatusDiProses
	} else if c.QueryParam("status") == "3" {
		status = entity.OrderStatusSelesai
	} else if c.QueryParam("status") != "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Status invalid",
		})
	}

	allOrder, err := models.GetAllOrder(c, claims.UserId, "", status)

	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataOrder(c, utils.JSONResponseDataOrder{
		Code:        http.StatusOK,
		GetAllOrder: allOrder,
		Message:     "Berhasil",
	})
}

func GetAllOrderPenjual(c echo.Context) error {
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)

	var status string
	if c.QueryParam("status") == "0" {
		status = entity.OrderStatusCancel
	} else if c.QueryParam("status") == "1" {
		status = entity.OrderStatusDipesan
	} else if c.QueryParam("status") == "2" {
		status = entity.OrderStatusDiProses
	} else if c.QueryParam("status") == "3" {
		status = entity.OrderStatusSelesai
	} else if c.QueryParam("status") != "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Status invalid",
		})
	}

	allOrder, err := models.GetAllOrder(c, "", claims.IdKeluarga, status)

	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataOrder(c, utils.JSONResponseDataOrder{
		Code:        http.StatusOK,
		GetAllOrder: allOrder,
		Message:     "Berhasil",
	})
}

func GetOrderByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	ord, err := models.GetOrderByID(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataOrder(c, utils.JSONResponseDataOrder{
		Code:         http.StatusOK,
		GetOrderByID: ord,
		Message:      "Berhasil",
	})
}

func OrderProses(c echo.Context) error {

	ord, err := models.GetOrderByID(c, c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	if ord.Status != entity.OrderStatusDipesan {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Flow Order Salah",
		})
	}
	ord.Status = entity.OrderStatusDiProses

	_, err = models.UpdateOrderById(c, ord.Id, &ord)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.Response(c, utils.JSONResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
	})
}

func OrderCancel(c echo.Context) error {
	ord, err := models.GetOrderByID(c, c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	if ord.Status != entity.OrderStatusDipesan {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Flow Order Salah",
		})
	}
	ord.Status = entity.OrderStatusCancel

	if ord.Pembayaran.Jenis == "Saldo" {
		dompet, err := models.GetDompetKeluargaByID(c, "", ord.Pembayaran.IdKeluargaPembeli)
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		dompet.Jumlah = dompet.Jumlah + ord.Harga_total

		_, err = models.UpdateDompetKeluargaById(c, dompet.Id, &dompet)
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	}

	_, err = models.UpdateOrderById(c, ord.Id, &ord)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.Response(c, utils.JSONResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
	})
}

func OrderSelesai(c echo.Context) error {
	ord, err := models.GetOrderByID(c, c.Param("id"))
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	if ord.Status != entity.OrderStatusDiProses {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Flow Order Salah",
		})
	}
	ord.Status = entity.OrderStatusSelesai

	if ord.Pembayaran.Jenis == "Saldo" {
		dompet, err := models.GetDompetKeluargaByID(c, "", ord.Pembayaran.IdKeluargaPenjual)
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		dompet.Jumlah = dompet.Jumlah + ord.Harga_total

		_, err = models.UpdateDompetKeluargaById(c, dompet.Id, &dompet)
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	} else {
		ord.Pembayaran.Status = entity.PembayaranTerbayar
		_, err = models.UpdatePembayaranById(c, ord.Pembayaran.Id, &ord.Pembayaran)
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	}

	_, err = models.UpdateOrderById(c, ord.Id, &ord)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.Response(c, utils.JSONResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
	})
}

func SoftDeleteOrderById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	_, err := models.GetOrderByID(c, id)

	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	_, err = models.SoftDeleteOrderById(c, id)

	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.Response(c, utils.JSONResponse{
		Code:    http.StatusBadRequest,
		Message: "Berhasil",
	})
}
