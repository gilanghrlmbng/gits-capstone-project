package controllers

import (
	"math/rand"
	"net/http"
	"src/api/models"
	"src/entity"
	"src/utils"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

func CreateKeluarga(c echo.Context) error {
	// Pertama inisiasi variable dulu
	k := new(entity.Keluarga)

	// kemudian ini buat dapetin request body dari mobile
	if err := c.Bind(k); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// terus ini ada validasi buat ngecek inputan dari reqeust body udah sesuai apa belum
	if err := k.ValidateCreate(); err.Code > 0 {
		return utils.ResponseError(c, err)
	}

	//Ini buat generate ID
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	k.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	// Ini buat masukin isi dari created_at nya
	k.CreatedAt = time.Now()

	// Ini fungsi dari models buat create data ke database
	keluarga, err := models.CreateKeluarga(k)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// Return datanya
	return utils.ResponseData(c, utils.JSONResponseData{
		Code:    http.StatusCreated,
		Data:    keluarga,
		Message: "Berhasil",
	})
}

func GetKeluarga(c echo.Context) error {
	Search := c.QueryParam("search")

	allKeluarga, err := models.GetAllKeluarga(Search)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseData(c, utils.JSONResponseData{
		Code:    http.StatusOK,
		Data:    allKeluarga,
		Message: "Berhasil",
	})
}

func GetKeluargaByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	k, err := models.GetKeluargaByID(id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseData(c, utils.JSONResponseData{
		Code:    http.StatusOK,
		Data:    k,
		Message: "Berhasil",
	})
}

func UpdateKeluargaById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	k := new(entity.Keluarga)

	if err := c.Bind(k); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_, err := models.GetKeluargaByID(id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.UpdateKeluargaById(id, k)
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

func SoftDeleteKeluargaById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	_, err := models.GetKeluargaByID(id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeleteKeluargaById(id)
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
