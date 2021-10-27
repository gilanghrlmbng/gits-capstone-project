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

	// Ini fungsi dari models buat create data ke database
	Rt, err := models.CreateRT(rt)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// Return datanya
	return utils.ResponseData(c, utils.JSONResponseData{
		Code:    http.StatusCreated,
		Data:    Rt,
		Message: "Berhasil",
	})
}

func GetAllRT(c echo.Context) error {
	allRT, err := models.GetAllRT()
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseData(c, utils.JSONResponseData{
		Code:    http.StatusOK,
		Data:    allRT,
		Message: "Berhasil",
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

	rt, err := models.GetRTByID(id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseData(c, utils.JSONResponseData{
		Code:    http.StatusOK,
		Data:    rt,
		Message: "Berhasil",
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

	_, err := models.GetRTByID(id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	rt.UpdatedAt = time.Now()

	_, err = models.UpdateRTById(id, rt)
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

	_, err := models.GetRTByID(id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeleteRTById(id)
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
