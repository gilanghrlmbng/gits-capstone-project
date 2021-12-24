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

	_, err := models.CreateOrder(c, ord)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.CreateBatchItemOrder(c, items)
	if err != nil {
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

func GetAllOrder(c echo.Context) error {
	allOrder, err := models.GetAllOrder(c)

	if err != nil {
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

func UpdateOrderById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "ID tidak valid",
		})
	}

	ord := new(entity.Order)

	if err := c.Bind(ord); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_, err := models.GetOrderByID(c, id)

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	ord.UpdatedAt = time.Now()

	_, err = models.UpdateOrderById(c, id, ord)

	if err != nil {
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
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	_, err = models.SoftDeleteOrderById(c, id)

	if err != nil {
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
