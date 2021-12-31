package models

import (
	"errors"
	"math/rand"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
)

func CreateForgetPasswordWarga(c echo.Context, p *entity.ForgetPasswordWarga) (entity.ForgetPasswordWarga, error) {
	db := db.GetDB(c)

	err := db.Create(&p)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.ForgetPasswordWarga{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.ForgetPasswordWarga{}, errors.New("gagal forget password")
	}

	return *p, nil
}

func GetWargaByForgetPasswordKode(c echo.Context, kode string) (entity.Warga, error) {
	var w []entity.Warga
	db := db.GetDB(c)

	err := db.Joins("ForgetPasswordWarga", db.Where(&entity.ForgetPasswordWarga{Kode: kode})).Find(&w)
	if err.Error != nil {
		c.Logger().Error(err.Error)
		return entity.Warga{}, errors.New("akun tidak ditemukan atau kode tidak valid")
	}
	var output entity.Warga
	for _, warga := range w {
		if warga.ForgetPasswordWarga != nil && warga.ForgetPasswordWarga.Kode == kode {
			output = warga
		}
	}
	if output.ForgetPasswordWarga.Kode != kode {
		c.Logger().Error(err.Error)
		return entity.Warga{}, errors.New("akun tidak ditemukan atau kode tidak valid")
	}

	return output, nil
}

func DeleteForgetPasswordWarga(c echo.Context, kode string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("kode = ?", kode).Delete(&entity.ForgetPasswordWarga{})

	if err.Error != nil || err.RowsAffected == 0 {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func GenerateKodeForgetPasswordWarga(c echo.Context, n int16) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	kode := string(b)
	db := db.GetDB(c)

	err := db.Where("kode = ?", kode).First(&entity.ForgetPasswordWarga{})
	for err.Error == nil {
		b = make([]byte, n)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		kode = string(b)
		err = db.Where("kode = ?", kode).First(&entity.ForgetPasswordWarga{})
	}
	return kode
}
