package models

import (
	"errors"
	"math/rand"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
)

func CreateRT(c echo.Context, rt *entity.Rt) (entity.Rt, error) {
	db := db.GetDB(c)

	err := db.Create(&rt)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Rt{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.Rt{}, errors.New("gagal membuat keluarga")
	}

	return *rt, nil
}

func GetAllRT(c echo.Context) ([]entity.Rt, error) {
	var RTs []entity.Rt
	db := db.GetDB(c)

	err := db.Find(&RTs)
	if err.Error != nil {
		c.Logger().Error(err)
		return RTs, err.Error
	}

	return RTs, nil
}

func GetAllRTWithPengurus(c echo.Context) ([]entity.Rt, error) {
	var RTs []entity.Rt
	db := db.GetDB(c)

	err := db.Preload("PengurusRT").Find(&RTs)
	if err.Error != nil {
		c.Logger().Error(err)
		return RTs, err.Error
	}

	return RTs, nil
}

func GetAllRTWithKeluarga(c echo.Context) ([]entity.Rt, error) {
	var RTs []entity.Rt
	db := db.GetDB(c)

	err := db.Preload("Keluarga").Find(&RTs)
	if err.Error != nil {
		c.Logger().Error(err)
		return RTs, err.Error
	}

	return RTs, nil
}

func GetRTByID(c echo.Context, id string) (entity.Rt, error) {
	var rt entity.Rt
	db := db.GetDB(c)

	err := db.First(&rt, "id = ?", id)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Rt{}, errors.New("id tidak ditemukan atau tidak valid")
	}

	return rt, nil
}

func UpdateRTById(c echo.Context, id string, rt *entity.Rt) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.Rt{}).Where("id = ?", id).Updates(rt)
	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeleteRTById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Rt{})
	if err.Error != nil || err.RowsAffected == 0 {
		c.Logger().Error(err)
		return 0, err.Error
	}

	return err.RowsAffected, nil
}

func GetRTByKode(c echo.Context, kode string) (entity.Rt, error) {
	var rt entity.Rt

	db := db.GetDB(c)

	err := db.First(&rt, "kode_rt = ?", kode)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Rt{}, errors.New("kode tidak ditemukan atau tidak valid")
	}

	return rt, nil
}

func GenerateKodeRT(c echo.Context, n int16) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	var rt entity.Rt
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	kode := string(b)
	db := db.GetDB(c)

	err := db.First(&rt, "kode_rt = ?", kode).Error
	if err == nil {
		for err == nil {
			b = make([]byte, n)
			for i := range b {
				b[i] = letterBytes[rand.Intn(len(letterBytes))]
			}
			kode = string(b)
			err = db.Where("kode_rt = ?", kode).First(&entity.Rt{}).Error
		}
	}

	return kode
}
