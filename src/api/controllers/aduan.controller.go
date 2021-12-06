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

func CreateAduan(c echo.Context) error {
	a := new(entity.Aduan)

	if err := c.Bind(a); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)

	warga, err := models.GetWargaByEmail(c, claims.Email)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	a.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	a.IdWarga = warga.Id
	a.CreatedBy = claims.Nama
	a.CreatedAt = time.Now()

	Aduan, err := models.CreateAduan(c, a)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusCreated,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataAduan(c, utils.JSONResponseDataAduan{
		Code:        http.StatusCreated,
		CreateAduan: Aduan,
		Message:     "Berhasil",
	})
}

func GetAllAduan(c echo.Context) error {
	allAduan, err := models.GetAllAduan(c)

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataAduan(c, utils.JSONResponseDataAduan{
		Code:        http.StatusOK,
		GetAllAduan: allAduan,
		Message:     "Berhasil",
	})
}

func GetAduanByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	Aduan, err := models.GetAduanByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataAduan(c, utils.JSONResponseDataAduan{
		Code:         http.StatusOK,
		GetAduanByID: Aduan,
		Message:      "Berhasil",
	})
}

func UpdateAduanById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	a := new(entity.Aduan)

	if err := c.Bind(a); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_, err := models.GetAduanByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	a.UpdatedAt = time.Now()

	_, err = models.UpdateAduanById(c, id, a)
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

func SoftDeleteAduanById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	_, err := models.GetAduanByID(c, id)

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeleteAduanById(c, id)

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
