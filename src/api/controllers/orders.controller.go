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
	ord := new(entity.Order)

	if err := c.Bind(ord); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)

	if claims.UserId == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}
	keluarga, err := models.GetWargaByEmail(c, claims.Email)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	w := keluarga.Warga[0]

	ord.IdWarga = w.Id

	// id pembayaran belum.
	// ord.IdPembayaran =

	if err := ord.ValidateCreate(); err.Code > 0 {
		return utils.ResponseError(c, err)
	}

	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	ord.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	ord.CreatedAt = time.Now()
	Order, err := models.CreateOrder(c, ord)

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataOrder(c, utils.JSONResponseDataOrder{
		Code:        http.StatusCreated,
		CreateOrder: Order,
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
