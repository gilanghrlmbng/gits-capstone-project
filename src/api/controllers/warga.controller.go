package controllers

import (
	"fmt"
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
	w.Gambar = fmt.Sprintf("https://dummyimage.com/500x500/29493B/fff&text=%c", w.Nama[0])
	// terus ini ada validasi buat ngecek inputan dari reqeust body udah sesuai apa belum
	if err := w.ValidateCreate(); err.Code > 0 {
		return utils.ResponseError(c, err)
	}

	k, err := models.GetKeluargaByKode(c, w.KodeKeluarga)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	w.IdKeluarga = k.Id

	cek, _ := models.GetWargaByEmail(c, w.Email)
	if cek.Id != "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Email sudah terdaftar",
		})
	}

	//Ini buat generate ID
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	w.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()

	pass := w.Password
	// Ini buat masukin isi dari created_at nya
	w.CreatedAt = time.Now()
	w.Password = utils.HashPassword(w.Password, w.Id)

	// Ini fungsi dari models buat create data ke database
	_, err = models.CreateWarga(c, w)
	if err != nil {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	// Return datanya
	return loginWarga(c, pass, k.IdRT, w)
}

func GetAllWarga(c echo.Context) error {
	allWarga, err := models.GetAllWarga(c, c.QueryParam("id_keluarga"))
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
	var id string
	paramid := c.Param("id")
	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)

	if paramid != "" {
		id = paramid
	} else if claims.UserId != "" {
		id = claims.UserId
	} else {
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
	if w.Password != "" {
		w.Password = utils.HashPassword(w.Password, w.Email)
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

func LoginWarga(c echo.Context) error {
	w := new(entity.Warga)

	if err := c.Bind(w); err != nil {
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	keluarga, err := models.GetWargaByEmail(c, w.Email)
	warga := keluarga.Warga[0]

	if err != nil {
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	isValid := utils.CheckPassword(w.Password, warga.Id, warga.Password)
	if !isValid {
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Message: "Password yang anda masukkan salah",
		})
	}
	token, err := utils.GenerateTokenWarga(c, warga.Nama, warga.Email, warga.Id, warga.IdKeluarga, keluarga.IdRT, utils.JWTStandartClaims)
	if err != nil {
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return utils.ResponseLogin(c, utils.JSONResponseLogin{
		Code:    http.StatusOK,
		Token:   token,
		Message: "Berhasil",
	})

}

func loginWarga(c echo.Context, pass, id_rt string, w *entity.Warga) error {

	isValid := utils.CheckPassword(pass, w.Id, w.Password)
	if !isValid {
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Token:   "",
			Message: "Password yang anda masukkan salah",
		})
	}
	token, err := utils.GenerateTokenWarga(c, w.Nama, w.Email, w.Id, w.IdKeluarga, id_rt, utils.JWTStandartClaims)
	if err != nil {
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Token:   "",
			Message: err.Error(),
		})
	}

	return utils.ResponseLogin(c, utils.JSONResponseLogin{
		Code:    http.StatusOK,
		Token:   token,
		Message: "Berhasil",
	})

}
