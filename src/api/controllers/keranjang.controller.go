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

func GetKeranjangByIDWarga(c echo.Context) error {
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	if claims.User == "pengurus" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}
	var id string
	if c.Param("id") != "" {
		id = c.Param("id")
	} else {
		id = claims.UserId
	}

	k, err := models.GetKeranjangByIDWarga(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataKeranjang(c, utils.JSONResponseDataKeranjang{
		Code:             http.StatusOK,
		GetKeranjangByID: k,
		Message:          "Berhasil",
	})
}

func UpdateKeranjangByIdWarga(c echo.Context) error {
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	if claims.User == "pengurus" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	var id string
	if c.Param("id") != "" {
		id = c.Param("id")
	} else {
		id = claims.UserId
	}

	k, err := models.GetKeranjangByIDWarga(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var itemsKeranjang []entity.ItemKeranjang

	if err := c.Bind(&itemsKeranjang); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	for _, item := range k.ItemKeranjang {
		err = models.HardDeleteItemKeranjang(c, item.Id)
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	}
	k.Harga_total = 0

	k.IdKeluargaPenjual = ""
	for idx, item := range itemsKeranjang {
		produk, err := models.GetProdukByID(c, item.IdProduk)
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		if k.IdKeluargaPenjual == "" {
			k.IdKeluargaPenjual = produk.IdKeluarga
		} else if produk.IdKeluarga != k.IdKeluargaPenjual {
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusBadRequest,
				Message: "Item yang dipesan harus dari toko yang sama",
			})
		}

		if err := item.ValidateCreate(); err.Code > 0 {
			return utils.ResponseError(c, err)
		}

		itemsKeranjang[idx].HargaTotal = produk.Harga * item.Jumlah
		k.Harga_total += itemsKeranjang[idx].HargaTotal

		entropys := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
		itemsKeranjang[idx].Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropys).String()

		itemsKeranjang[idx].IdKeranjang = k.Id
		itemsKeranjang[idx].CreatedAt = time.Now()
	}

	_, err = models.UpdateKeranjangById(c, k.Id, &k)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	if len(itemsKeranjang) > 0 {
		_, err = models.CreateBatchItemKeranjang(c, itemsKeranjang)
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	}

	return utils.ResponseDataKeranjang(c, utils.JSONResponseDataKeranjang{
		Code:    http.StatusOK,
		Message: "Berhasil",
	})
}

func TambahItemKeranjang(c echo.Context) error {
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	if claims.User == "pengurus" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	var id string
	if c.Param("id") != "" {
		id = c.Param("id")
	} else {
		id = claims.UserId
	}

	k, err := models.GetKeranjangByIDWarga(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var itemKeranjang entity.ItemKeranjang

	if err := c.Bind(&itemKeranjang); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var sama bool = false
	var hargaSatuan int64 = 0
	for _, item := range k.ItemKeranjang {
		if item.IdProduk == itemKeranjang.IdProduk {
			hargaSatuan = item.HargaTotal / item.Jumlah
			item.Jumlah = item.Jumlah + 1
			item.HargaTotal = item.Jumlah * hargaSatuan
			item.UpdatedAt = time.Now()
			_, err := models.UpdateItemKeranjangByID(c, item.Id, item)
			if err != nil {
				c.Logger().Error(err)
				return utils.ResponseError(c, utils.Error{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				})
			}
			sama = true
			break
		}
	}

	if sama {
		k.Harga_total += hargaSatuan

		_, err = models.UpdateKeranjangById(c, k.Id, &k)
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}

		c.Logger().Info("Masuk nambah")
		return utils.ResponseDataKeranjang(c, utils.JSONResponseDataKeranjang{
			Code:    http.StatusOK,
			Message: "Berhasil",
		})
	}

	entropys := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	itemKeranjang.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropys).String()

	produk, err := models.GetProdukByID(c, itemKeranjang.IdProduk)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if k.IdKeluargaPenjual == "" {
		k.IdKeluargaPenjual = produk.IdKeluarga
	} else if produk.IdKeluarga != k.IdKeluargaPenjual {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Item yang dipesan harus dari toko yang sama",
		})
	}

	itemKeranjang.IdKeranjang = k.Id
	itemKeranjang.CreatedAt = time.Now()

	itemKeranjang.HargaTotal = produk.Harga * itemKeranjang.Jumlah
	k.Harga_total += itemKeranjang.HargaTotal

	_, err = models.UpdateKeranjangById(c, k.Id, &k)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.CreateItemKeranjang(c, &itemKeranjang)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataKeranjang(c, utils.JSONResponseDataKeranjang{
		Code:    http.StatusOK,
		Message: "Berhasil",
	})
}

func TambahQuantityItemKeranjang(c echo.Context) error {
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	if claims.User == "pengurus" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	k, err := models.GetKeranjangByIDWarga(c, claims.UserId)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var itemKeranjang entity.ItemKeranjang

	if err := c.Bind(&itemKeranjang); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var hargaSatuan int64 = 0
	for _, item := range k.ItemKeranjang {
		if item.IdProduk == c.Param("id") {
			hargaSatuan = item.HargaTotal / item.Jumlah
			item.Jumlah += item.Jumlah
			item.HargaTotal = item.Jumlah * hargaSatuan
			item.UpdatedAt = time.Now()
			_, err := models.UpdateItemKeranjangByID(c, item.Id, item)
			if err != nil {
				c.Logger().Error(err)
				return utils.ResponseError(c, utils.Error{
					Code:    http.StatusInternalServerError,
					Message: err.Error(),
				})
			}
			break
		}
	}

	k.Harga_total += hargaSatuan

	_, err = models.UpdateKeranjangById(c, k.Id, &k)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	c.Logger().Info("Masuk nambah")
	return utils.ResponseDataKeranjang(c, utils.JSONResponseDataKeranjang{
		Code:    http.StatusOK,
		Message: "Berhasil",
	})
}

func KurangQuantityItemKeranjang(c echo.Context) error {
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	if claims.User == "pengurus" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	k, err := models.GetKeranjangByIDWarga(c, claims.UserId)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var itemKeranjang entity.ItemKeranjang

	if err := c.Bind(&itemKeranjang); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var hargaSatuan int64 = 0
	for _, item := range k.ItemKeranjang {
		if item.IdProduk == c.Param("id") {
			if item.Jumlah == 1 {

			} else {
				hargaSatuan = item.HargaTotal / item.Jumlah
				item.Jumlah -= 1
				item.HargaTotal = item.Jumlah * hargaSatuan
				item.UpdatedAt = time.Now()
				_, err := models.UpdateItemKeranjangByID(c, item.Id, item)
				if err != nil {
					c.Logger().Error(err)
					return utils.ResponseError(c, utils.Error{
						Code:    http.StatusInternalServerError,
						Message: err.Error(),
					})
				}
			}
			break
		}
	}

	k.Harga_total -= hargaSatuan

	_, err = models.UpdateKeranjangById(c, k.Id, &k)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	c.Logger().Info("Masuk kurang")
	return utils.ResponseDataKeranjang(c, utils.JSONResponseDataKeranjang{
		Code:    http.StatusOK,
		Message: "Berhasil",
	})
}
