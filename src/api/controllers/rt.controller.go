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

func CreateRT(c echo.Context) error {
	// Pertama inisiasi variable dulu
	rt := new(entity.Rt)

	// kemudian ini buat dapetin request body dari mobile
	if err := c.Bind(rt); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	// terus ini ada validasi buat ngecek inputan dari reqeust body udah sesuai apa belum
	if err := rt.ValidateCreate(); err.Code > 0 {
		return utils.ResponseError(c, err)
	}

	//Ini buat generate ID
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	rt.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	// Ini buat masukin isi dari created_at nya
	rt.CreatedAt = time.Now()
	rt.KodeRT = models.GenerateKodeRT(c, 6)

	// Ini fungsi dari models buat create data ke database
	Rt, err := models.CreateRT(c, rt)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	eror := CreateDompet(c, rt.Id)
	if eror.Code != http.StatusCreated {
		c.Logger().Error("Failed to Create Dompet RT")
		return utils.ResponseError(c, eror)
	}

	// Return datanya
	return utils.ResponseDataRT(c, utils.JSONResponseDataRT{
		Code:     http.StatusCreated,
		CreateRT: Rt,
		Message:  "Berhasil",
	})
}

func GetAllRT(c echo.Context) error {
	allRT, err := models.GetAllRT(c)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataRT(c, utils.JSONResponseDataRT{
		Code:     http.StatusOK,
		GetAllRT: allRT,
		Message:  "Berhasil",
	})
}

func GetAllRTWithPengurus(c echo.Context) error {
	allRT, err := models.GetAllRTWithPengurus(c)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataRT(c, utils.JSONResponseDataRT{
		Code:             http.StatusOK,
		GetAllRTPengurus: allRT,
		Message:          "Berhasil",
	})
}

func GetAllRTWithKeluarga(c echo.Context) error {
	allRT, err := models.GetAllRTWithKeluarga(c)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataRT(c, utils.JSONResponseDataRT{
		Code:             http.StatusOK,
		GetAllRTKeluarga: allRT,
		Message:          "Berhasil",
	})
}

func GetRTByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	rt, err := models.GetRTByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataRT(c, utils.JSONResponseDataRT{
		Code:      http.StatusOK,
		GetRTByID: rt,
		Message:   "Berhasil",
	})
}

func UpdateRTById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	rt := new(entity.Rt)

	if err := c.Bind(rt); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_, err := models.GetRTByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	rt.UpdatedAt = time.Now()

	_, err = models.UpdateRTById(c, id, rt)
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

func SoftDeleteRTById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	_, err := models.GetRTByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeleteRTById(c, id)
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
