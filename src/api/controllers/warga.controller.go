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

func CreateWarga(c echo.Context) error {
	// Pertama inisiasi variable dulu
	w := new(entity.Warga)

	// kemudian ini buat dapetin request body dari mobile
	if err := c.Bind(w); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// terus ini ada validasi buat ngecek inputan dari reqeust body udah sesuai apa belum
	if err := w.ValidateCreate(); err.Code > 0 {
		return utils.ResponseError(c, err)
	}

	_, err := models.GetWargaByEmail(c, w.Email)
	if err.Error() != "email tidak ditemukan" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	//Ini buat generate ID
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	w.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	// Ini buat masukin isi dari created_at nya
	w.CreatedAt = time.Now()

	// Ini fungsi dari models buat create data ke database
	W, err := models.CreateWarga(c, w)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// Return datanya
	return utils.ResponseDataWarga(c, utils.JSONResponseDataWarga{
		Code:        http.StatusCreated,
		CreateWarga: W,
		Message:     "Berhasil",
	})
}

func GetAllWarga(c echo.Context) error {
	allWarga, err := models.GetAllWarga(c)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataWarga(c, utils.JSONResponseDataWarga{
		Code:        http.StatusOK,
		GetAllWarga: allWarga,
		Message:     "Berhasil",
	})
}

func GetWargaByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	w, err := models.GetWargaByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataWarga(c, utils.JSONResponseDataWarga{
		Code:         http.StatusOK,
		GetWargaByID: w,
		Message:      "Berhasil",
	})
}

func UpdateWargaById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	w := new(entity.Warga)

	if err := c.Bind(w); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_, err := models.GetWargaByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	w.UpdatedAt = time.Now()

	_, err = models.UpdateWargaById(c, id, w)
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

func SoftDeleteWargaById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	_, err := models.GetWargaByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeleteWargaById(c, id)
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
