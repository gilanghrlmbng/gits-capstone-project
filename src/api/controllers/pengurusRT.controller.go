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

func CreatePengurus(c echo.Context) error {
	prt := new(entity.PengurusRT)

	if err := c.Bind(prt); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := prt.ValidateCreate(); err.Code > 0 {
		return utils.ResponseError(c, err)
	}

	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	prt.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	prt.CreatedAt = time.Now()

	PengurusRT, err := models.CreatePengurusRT(c, prt)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataPengurusRT(c, utils.JSONResponseDataPengurusRT{
		Code:           http.StatusCreated,
		CreatePengurus: PengurusRT,
		Message:        "Berhasil",
	})
}

func GetAllPengurusRT(c echo.Context) error {
	allPengurusRT, err := models.GetAllPengurusRT(c)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataPengurusRT(c, utils.JSONResponseDataPengurusRT{
		Code:           http.StatusOK,
		GetAllPengurus: allPengurusRT,
		Message:        "Berhasil",
	})
}

func GetPengurusByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	prt, err := models.GetPengurusByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataPengurusRT(c, utils.JSONResponseDataPengurusRT{
		Code:            http.StatusOK,
		GetPengurusByID: prt,
		Message:         "Berhasil",
	})
}

func UpdatePengurusById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	prt := new(entity.PengurusRT)

	if err := c.Bind(prt); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_, err := models.GetPengurusByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	prt.UpdatedAt = time.Now()

	_, err = models.UpdatePengurusById(c, id, prt)
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

func SoftDeletePengurusById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id todak valid",
		})
	}

	_, err := models.GetPengurusByID(c, id)

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeletePengurusById(c, id)

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

func LoginPengurus(c echo.Context) error {
	prt := new(entity.PengurusRT)

	if err := c.Bind(prt); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	pengurus, err := models.PengurusSearchEmail(c, prt.Email)

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	passTrue := utils.CheckPassword(prt.Password, prt.Email, pengurus.Password)

	if !passTrue {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Password Salah",
		})
	}

	token, err := utils.GenerateTokenPengurus(c, pengurus.Nama, pengurus.Email, pengurus.Id, pengurus.IdRT, utils.JWTStandartClaims)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return utils.ResponseData(c, utils.JSONResponseData{
		Code:    http.StatusCreated,
		Data:    map[string]string{"token": token, "nama": pengurus.Nama, "email": pengurus.Email},
		Message: "Berhasil",
	})
}
