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
	item := new(entity.ItemOrder)
	ord := new(entity.Order)

	if err := c.Bind(item); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	produk, err := models.GetProdukByID(c, item.IdProduk)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	item.HargaTotal = produk.Harga * item.Jumlah

	if err := item.ValidateCreate(); err.Code > 0 {
		return utils.ResponseError(c, err)
	}

	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)

	if claims.UserId == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	warga, err := models.GetWargaByEmail(c, claims.Email)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	orders, err := models.GetAllOrder(c)

	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)

	ord.IdWarga = warga.Id

	if len(orders) == 0 {
		ord.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
		_, err = models.CreateOrder(c, ord)
	} else {
		if orders[len(orders)-1].IdWarga == ord.IdWarga {
			ord.Id = orders[len(orders)-1].Id
		} else {
			ord.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
			_, err = models.CreateOrder(c, ord)
		}
	}

	item.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	item.IdOrder = ord.Id

	_, err = models.CreateItemOrder(c, item)

	items, err := models.GetAllItemOrder(c)

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	for i := 0; i < len(items); i++ {
		if items[i].IdOrder == item.IdOrder {
			ord.Harga_total = ord.Harga_total + items[i].HargaTotal
			_, err = models.UpdateOrderById(c, ord.Id, ord)
		}

	}

	ord.CreatedAt = time.Now()

	allOrder, err := models.GetAllOrder(c)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataOrder(c, utils.JSONResponseDataOrder{
		Code:        http.StatusCreated,
		CreateOrder: allOrder,
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
