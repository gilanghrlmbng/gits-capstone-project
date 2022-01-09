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

func CreateInformasi(c echo.Context) error {
	p := new(entity.Informasi)

	if err := c.Bind(p); err != nil {
		c.Logger().Error(err)
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

	p.IdRT = claims.IdRT
	p.CreatedBy = claims.Nama
	if err := p.ValidateCreate(); err.Code > 0 {
		c.Logger().Error(err)
		return utils.ResponseError(c, err)
	}

	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	p.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	p.CreatedAt = time.Now()

	Informasi, err := models.CreateInformasi(c, p)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataInformasi(c, utils.JSONResponseDataInformasi{
		Code:            http.StatusCreated,
		CreateInformasi: Informasi,
		Message:         "Berhasil",
	})
}

func GetAllInformasi(c echo.Context) error {

	allInformasi, err := models.GetAllInformasi(c, c.QueryParam("id_keluarga"))
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataInformasi(c, utils.JSONResponseDataInformasi{
		Code:            http.StatusOK,
		GetAllInformasi: allInformasi,
		Message:         "Berhasil",
	})
}

func GetInformasiByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	p, err := models.GetInformasiByID(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataInformasi(c, utils.JSONResponseDataInformasi{
		Code:             http.StatusOK,
		GetInformasiByID: p,
		Message:          "Berhasil",
	})
}

func GetInfoTerkini(c echo.Context) error {
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	if claims.IdRT == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	p, err := models.GetAllInformasiByKategori(c, claims.IdRT, "Informasi")
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataInformasi(c, utils.JSONResponseDataInformasi{
		Code:            http.StatusOK,
		GetAllInformasi: p,
		Message:         "Berhasil",
	})
}

func GetKegiatanWarga(c echo.Context) error {
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	if claims.IdRT == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	p, err := models.GetAllInformasiByKategori(c, claims.IdRT, "Kegiatan")
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataInformasi(c, utils.JSONResponseDataInformasi{
		Code:            http.StatusOK,
		GetAllInformasi: p,
		Message:         "Berhasil",
	})
}

func UpdateInformasiById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	p := new(entity.Informasi)

	if err := c.Bind(p); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_, err := models.GetInformasiByID(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	p.UpdatedAt = time.Now()

	_, err = models.UpdateInformasiById(c, id, p)
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

func SoftDeleteInformasiById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id todak valid",
		})
	}

	_, err := models.GetInformasiByID(c, id)

	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeleteInformasiById(c, id)

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
