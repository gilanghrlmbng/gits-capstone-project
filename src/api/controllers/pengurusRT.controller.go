package controllers

import (
	"fmt"
	"math/rand"
	"net/http"
	"src/api/models"
	"src/entity"
	"src/utils"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

func CreatePengurus(c echo.Context) error {
	prt := new(entity.PengurusRT)

	if err := c.Bind(prt); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	prt.Gambar = "default_image"

	if strings.HasPrefix(prt.NoHandphone, "62") {
		prt.NoHandphone = fmt.Sprintf("0%s", strings.SplitAfter(prt.NoHandphone, "62")[1])
	}
	if strings.HasPrefix(prt.NoHandphone, "+62") {
		prt.NoHandphone = fmt.Sprintf("0%s", strings.SplitAfter(prt.NoHandphone, "+62")[1])
	}

	if err := prt.ValidateCreate(); err.Code > 0 {
		c.Logger().Error(err)
		return utils.ResponseError(c, err)
	}

	rt, err := models.GetRTByKode(c, prt.KodeRT)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	prt.IdRT = rt.Id

	cek, err := models.PengurusSearchEmail(c, prt.Email)
	if err != nil && err.Error() != "email tidak ditemukan atau tidak valid" {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	if cek.Id != "" {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Email sudah terdaftar",
		})
	}

	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	prt.Id = ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	pass := prt.Password
	prt.CreatedAt = time.Now()
	prt.Password = utils.HashPassword(prt.Password, prt.Id)
	_, err = models.CreatePengurusRT(c, prt)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return loginPengurus(c, pass, prt)
}

func GetAllPengurusRT(c echo.Context) error {
	allPengurusRT, err := models.GetAllPengurusRT(c)
	if err != nil {
		c.Logger().Error(err)
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

	prt, err := models.GetPengurusByID(c, id)
	if err != nil {
		c.Logger().Error(err)
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
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if err := prt.ValidateUpdate(); err.Code > 0 {
		c.Logger().Error(err)
		return utils.ResponseError(c, err)
	}

	_, err := models.GetPengurusByID(c, id)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	if prt.Password != "" {
		prt.Password = utils.HashPassword(prt.Password, prt.Email)
	}
	prt.UpdatedAt = time.Now()

	_, err = models.UpdatePengurusById(c, id, prt)
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

func LoginPengurus(c echo.Context) error {
	prt := new(entity.PengurusRT)

	if err := c.Bind(prt); err != nil {
		c.Logger().Error(err)
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	pengurus, err := models.PengurusSearchEmail(c, prt.Email)

	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	passTrue := utils.CheckPassword(prt.Password, pengurus.Id, pengurus.Password)

	if !passTrue {
		return utils.ResponseErrorLogin(c, utils.ErrorLogin{
			Code:    http.StatusBadRequest,
			Message: "Password Salah",
		})
	}

	if prt.TokenFirebase != "" {
		_, err := models.UpdatePengurusById(c, pengurus.Id, &entity.PengurusRT{TokenFirebase: prt.TokenFirebase})
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseErrorLogin(c, utils.ErrorLogin{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	}

	token, err := utils.GenerateTokenPengurus(c, pengurus.Nama, pengurus.Email, pengurus.Id, pengurus.IdRT, utils.ExpiredHour)
	if err != nil {
		c.Logger().Error(err)
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

func loginPengurus(c echo.Context, pass string, prt *entity.PengurusRT) error {

	passTrue := utils.CheckPassword(pass, prt.Id, prt.Password)

	if !passTrue {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Password Salah",
		})
	}

	if prt.TokenFirebase != "" {
		_, err := models.UpdatePengurusById(c, prt.Id, &entity.PengurusRT{TokenFirebase: prt.TokenFirebase})
		if err != nil {
			c.Logger().Error(err)
			return utils.ResponseErrorLogin(c, utils.ErrorLogin{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
			})
		}
	}

	token, err := utils.GenerateTokenPengurus(c, prt.Nama, prt.Email, prt.Id, prt.IdRT, utils.ExpiredHour)
	if err != nil {
		c.Logger().Error(err)
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

func ForgetPasswordPengurus(c echo.Context) error {
	fp := new(ForgetPasswordRequest)

	if err := c.Bind(fp); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	p, err := models.PengurusSearchEmail(c, fp.Email)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	entropy := ulid.Monotonic(rand.New(rand.NewSource(time.Now().UnixNano())), 0)
	fpId := ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
	kode := models.GenerateKodeForgetPasswordPengurus(c, 6)

	forgetPass := entity.ForgetPasswordPengurus{
		Id:         fpId,
		IdPengurus: p.Id,
		Kode:       kode,
		CreatedAt:  time.Now(),
	}

	fpw, err := models.CreateForgetPasswordPengurus(c, &forgetPass)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	p.ForgetPasswordPengurus = &fpw
	_, err = models.UpdatePengurusById(c, p.Id, &p)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	err = utils.SendEmail(c, fp.Email, "Kode Reset Password", fmt.Sprintf("Berikut ini adalah kode Verifikasi untuk reset password akun pengurus anda <br><br> Kode: <b>%s</b> <br><br> abaikan jika anda tidak sedang mereset password", kode))
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

func ResetPasswordPengurusByKode(c echo.Context) error {
	rp := new(ResetPasswordRequest)

	if err := c.Bind(rp); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if !utils.CheckStrengthPassword(rp.Password) {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Password panjangnya min. 8 karakter, serta mengandung min. 1 huruf besar, 1 huruf kecil, dan 1 angka!",
		})
	}

	p, err := models.GetPengurusByForgetPasswordKode(c, rp.Kode)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	p.ForgetPasswordPengurus = &entity.ForgetPasswordPengurus{}
	p.Password = utils.HashPassword(rp.Password, p.Id)

	_, err = models.UpdatePengurusById(c, p.Id, &p)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	_, err = models.DeleteForgetPasswordPengurus(c, rp.Kode)
	if err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	c.Logger().Info(p)
	return utils.Response(c, utils.JSONResponse{
		Code:    http.StatusOK,
		Message: "Berhasil",
	})
}

func GantiPasswordPengurus(c echo.Context) error {
	cp := new(ChangePasswordRequest)

	if err := c.Bind(cp); err != nil {
		c.Logger().Error(err)
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	if !utils.CheckStrengthPassword(cp.NewPaswword) {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Password panjangnya min. 8 karakter, serta mengandung min. 1 huruf besar, 1 huruf kecil, dan 1 angka!",
		})
	}

	userData := c.Get("user").(*jwt.Token)
	claims := userData.Claims.(*utils.JWTCustomClaims)

	p, _ := models.GetPengurusByID(c, claims.UserId)

	isValid := utils.CheckPassword(cp.Password, claims.UserId, p.Password)
	if !isValid {
		return utils.ResponseError(c, utils.Error{
			Code:    http.StatusBadRequest,
			Message: "Password yang anda masukkan salah",
		})
	}

	_, err := models.UpdatePengurusById(c, claims.UserId, &entity.PengurusRT{Password: utils.HashPassword(cp.NewPaswword, p.Id)})
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
