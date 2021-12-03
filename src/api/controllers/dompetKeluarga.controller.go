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

func CreateDompetKeluarga(c echo.Context, id_keluarga string) utils.Error {
	d := entity.DompetKeluarga{
		Jumlah: 0,
	}

	d.IdKeluarga = id_keluarga

	if err := d.ValidateCreate(); err.Code > 0 {
		return err
	}

	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	d.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	d.CreatedAt = time.Now()

	_, err := models.CreateDompetKeluarga(c, &d)
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

func GetAllDompetKeluarga(c echo.Context) error {
	allDompet, err := models.GetAllDompetKeluarga(c)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataDompetKeluarga(c, utils.JSONResponseDataDompetKeluarga{
		Code:                 http.StatusOK,
		GetAllDompetKeluarga: allDompet,
		Message:              "Berhasil",
	})
}

func GetDompetKeluargaByID(c echo.Context) error {
	var id_keluarga string
	id := c.Param("id")
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	c.Logger().Info(claims)
	if claims.IdKeluarga != "" && claims.User == "warga" {
		id_keluarga = claims.IdKeluarga
	} else if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	d, err := models.GetDompetKeluargaByID(c, id, id_keluarga)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataDompetKeluarga(c, utils.JSONResponseDataDompetKeluarga{
		Code:                  http.StatusOK,
		GetDompetKeluargaByID: d,
		Message:               "Berhasil",
	})
}

func UpdateDompetKeluargaById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	d := new(entity.DompetKeluarga)

	if err := c.Bind(d); err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_, err := models.GetDompetKeluargaByID(c, id, "")

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	d.UpdatedAt = time.Now()

	_, err = models.UpdateDompetKeluargaById(c, id, d)
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

func SoftDeleteDompetKeluargaById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	_, err := models.GetDompetKeluargaByID(c, id, "")

	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeleteDompetKeluargaById(c, id)

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
