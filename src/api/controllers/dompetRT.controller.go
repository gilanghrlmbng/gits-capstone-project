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

func CreateDompet(c echo.Context, id_rt string) utils.Error {
	d := entity.DompetRT{
		Jumlah: 0,
	}

	d.IdRT = id_rt

	if err := d.ValidateCreate(); err.Code > 0 {
		return err
	}

	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	d.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	d.CreatedAt = time.Now()

	_, err := models.CreateDompet(c, &d)
	if err != nil {
		return utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return utils.Error{
		Code:    http.StatusCreated,
		Message: "Berhasil",
	}
}

func GetAllDompet(c echo.Context) error {
	allDompet, err := models.GetAllDompet(c)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataDompet(c, utils.JSONResponseDataDompetRT{
		Code:         http.StatusOK,
		GetAllDompet: allDompet,
		Message:      "Berhasil",
	})
}

func GetDompetByID(c echo.Context) error {
	var id_rt string
	id := c.Param("id")
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	if claims.IdRT != "" && claims.User == "pengurus" {
		id_rt = claims.IdRT
	} else if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	d, err := models.GetDompetByID(c, id, id_rt)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataDompet(c, utils.JSONResponseDataDompetRT{
		Code:          http.StatusOK,
		GetDompetByID: d,
		Message:       "Berhasil",
	})
}

func TopUpDompetRT(c echo.Context) error {
	d := new(entity.DompetRT)

	if err := c.Bind(d); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	if claims.IdRT == "" || claims.User != "pengurus" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	dompet, err := models.GetDompetByID(c, "", claims.IdRT)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	dompet.Jumlah = dompet.Jumlah + d.Jumlah

	_, err = models.UpdateDompetById(c, dompet.Id, &dompet)
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

func WithdrawDompetRT(c echo.Context) error {
	d := new(entity.DompetRT)

	if err := c.Bind(d); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	if claims.IdRT == "" || claims.User != "pengurus" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	dompet, err := models.GetDompetByID(c, "", claims.IdRT)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	dompet.Jumlah = dompet.Jumlah - d.Jumlah

	_, err = models.UpdateDompetById(c, dompet.Id, &dompet)
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

func UpdateDompetById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	d := new(entity.DompetRT)

	if err := c.Bind(d); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_, err := models.GetDompetByID(c, id, "")

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	d.UpdatedAt = time.Now()

	_, err = models.UpdateDompetById(c, id, d)
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

func SoftDeleteDompetById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	_, err := models.GetDompetByID(c, id, "")

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeleteDompetById(c, id)

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
