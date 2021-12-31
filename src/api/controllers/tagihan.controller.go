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

func CreateTagihan(c echo.Context) error {
	t := new(entity.Tagihan)

	if err := c.Bind(t); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)
	if claims.IdRT == "" && claims.User != "pengurus" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	if err := t.ValidateCreate(); err.Code > 0 {
		c.Logger().Error(err)
		return utils.ResponseError(c, err)
	}

	Keluargas, err := models.GetAllKeluargaByRT(c, claims.IdRT)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	t.CreatedAt = time.Now()

	for _, kel := range Keluargas {
		entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
		Id := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

		tag := entity.Tagihan{
			Id:         Id,
			IdKeluarga: kel.Id,
			IdRT:       claims.IdRT,
			Nama:       t.Nama,
			Detail:     t.Detail,
			Jumlah:     t.Jumlah,
			CreatedAt:  t.CreatedAt,
		}

		_, err := models.CreateTagihan(c, &tag)
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseError(c, utils.Error{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	}

	return utils.ResponseDataTagihan(c, utils.JSONResponseDataTagihan{
		Code:          http.StatusCreated,
		CreateTagihan: t,
		Message:       "Berhasil",
	})
}

func GetAllTagihan(c echo.Context) error {
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)

	var allTagihan []entity.Tagihan
	var err error
	if claims.User == "pengurus" {
		allTagihan, err = models.GetAllTagihan(c, "", claims.IdRT)
	} else if claims.User == "warga" {
		allTagihan, err = models.GetAllTagihan(c, claims.IdKeluarga, "")
	} else {
		allTagihan, err = models.GetAllTagihan(c, "", "")
	}

	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return utils.ResponseDataTagihan(c, utils.JSONResponseDataTagihan{
		Code:          http.StatusOK,
		GetAllTagihan: allTagihan,
		Message:       "Berhasil",
	})
}

func GetTagihanByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	t, err := models.GetTagihanByID(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return utils.ResponseDataTagihan(c, utils.JSONResponseDataTagihan{
		Code:           http.StatusOK,
		GetTagihanByID: t,
		Message:        "Berhasil",
	})
}

func BayarTagihanByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	t, err := models.GetTagihanByID(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	if t.Terbayar == "true" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Tagihan ini sudah terbayar",
		})
	}

	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)

	if claims.IdKeluarga == "" && claims.User != "warga" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Maaf anda tidak memiliki akses ini",
		})
	}

	dompet, err := models.GetDompetKeluargaByID(c, "", claims.IdKeluarga)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	dompetRT, err := models.GetDompetByID(c, "", claims.IdRT)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	if dompet.Jumlah < t.Jumlah {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Saldo tidak mencukupi",
		})
	}

	dompet.Jumlah = dompet.Jumlah - t.Jumlah
	dompetRT.Jumlah = dompetRT.Jumlah + t.Jumlah
	t.Terbayar = "true"

	_, err = models.UpdateDompetKeluargaById(c, dompet.Id, &dompet)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.UpdateDompetById(c, dompetRT.Id, &dompetRT)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.UpdateTagihanById(c, t.Id, &t)
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

func UpdateTagihanById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id tidak valid",
		})
	}

	t := new(entity.Tagihan)

	if err := c.Bind(t); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	_, err := models.GetTagihanByID(c, id)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	t.UpdatedAt = time.Now()

	_, err = models.UpdateTagihanById(c, id, t)
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

func SoftDeleteTagihanById(c echo.Context) error {
	id := c.Param("id")

	if id == "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Id todak valid",
		})
	}

	_, err := models.GetTagihanByID(c, id)

	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.SoftDeleteTagihanById(c, id)

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
