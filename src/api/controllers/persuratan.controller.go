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

func CreatePersuratan(c echo.Context) error {
	s := new(entity.Persuratan)

	if err := c.Bind(s); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	if claims.IdRT == "" && claims.User != "warga" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	s.IdRT = claims.IdRT
	s.IdWarga = claims.UserId
	s.Status = entity.StatusPersuratanTerkirim

	if err := s.ValidateCreate(); err.Code > 0 {
		c.Logger().Error(err)
		return utils.ResponseError(c, err)
	}

	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	s.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	s.CreatedAt = time.Now()

	Persuratan, err := models.CreatePersuratan(c, s)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataPersuratan(c, utils.JSONResponseDataPersuratan{
		Code:             http.StatusCreated,
		CreatePersuratan: Persuratan,
		Message:          "Berhasil",
	})
}

func GetAllPersuratan(c echo.Context) error {
	var allPersuratan []entity.Persuratan
	var err error
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)

	var status string = ""
	if c.QueryParam("status") == "0" {
		status = entity.StatusPersuratanTolak
	} else if c.QueryParam("status") == "1" {
		status = entity.StatusPersuratanTerkirim
	} else if c.QueryParam("status") == "2" {
		status = entity.StatusPersuratanDiproses
	} else if c.QueryParam("status") == "3" {
		status = entity.StatusPersuratanSelesai
	}

	if claims.User == "pengurus" {
		allPersuratan, err = models.GetAllPersuratan(c, claims.IdRT, "", status)
	} else {
		allPersuratan, err = models.GetAllPersuratan(c, "", claims.UserId, status)
	}

	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataPersuratan(c, utils.JSONResponseDataPersuratan{
		Code:             http.StatusOK,
		GetAllPersuratan: allPersuratan,
		Message:          "Berhasil",
	})
}

func GetPersuratanByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	s, err := models.GetPersuratanByID(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataPersuratan(c, utils.JSONResponseDataPersuratan{
		Code:              http.StatusOK,
		GetPersuratanByID: s,
		Message:           "Berhasil",
	})
}

func UpdatePersuratanById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	Surat, err := models.GetPersuratanByID(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	if Surat.Status != entity.StatusPersuratanTerkirim {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: "Maaf surat sudah diproses, data tidak bisa diubah lagi",
		})
	}

	s := new(entity.Persuratan)

	if err := c.Bind(s); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	s.UpdatedAt = time.Now()

	_, err = models.UpdatePersuratanById(c, id, s)
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

func SoftDeletePersuratanById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id todak valid",
		})
	}

	_, err := models.GetPersuratanByID(c, id)

	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeletePersuratanById(c, id)

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
