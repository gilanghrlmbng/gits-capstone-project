package models

import (
	"errors"
	"math/rand"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
)

func CreateForgetPasswordPengurus(c echo.Context, p *entity.ForgetPasswordPengurus) (entity.ForgetPasswordPengurus, error) {
	db := db.GetDB(c)

	err := db.Create(&p)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.ForgetPasswordPengurus{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.ForgetPasswordPengurus{}, errors.New("gagal forget password")
	}

	return *p, nil
}

func GetPengurusByForgetPasswordKode(c echo.Context, kode string) (entity.PengurusRT, error) {
	var prt []entity.PengurusRT
	db := db.GetDB(c)

	err := db.Joins("ForgetPasswordPengurus", db.Where(&entity.ForgetPasswordPengurus{Kode: kode})).Find(&prt)
	if err.Error != nil {
		c.Logger().Error(err.Error)
		return entity.PengurusRT{}, errors.New("akun tidak ditemukan atau kode tidak valid")
	}
	var output entity.PengurusRT
	for _, p := range prt {
		if p.ForgetPasswordPengurus != nil && p.ForgetPasswordPengurus.Kode == kode {
			output = p
		}
	}
	if output.ForgetPasswordPengurus.Kode != kode {
		c.Logger().Error(err.Error)
		return entity.PengurusRT{}, errors.New("akun tidak ditemukan atau kode tidak valid")
	}

	return output, nil
}

func DeleteForgetPasswordPengurus(c echo.Context, kode string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("kode = ?", kode).Delete(&entity.ForgetPasswordPengurus{})

	if err.Error != nil || err.RowsAffected == 0 {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func GenerateKodeForgetPasswordPengurus(c echo.Context, n int16) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	kode := string(b)
	db := db.GetDB(c)

	err := db.Where("kode = ?", kode).First(&entity.ForgetPasswordPengurus{})
	for err.Error == nil {
		b = make([]byte, n)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		kode = string(b)
		err = db.Where("kode = ?", kode).First(&entity.ForgetPasswordPengurus{})
	}
	return kode
}
