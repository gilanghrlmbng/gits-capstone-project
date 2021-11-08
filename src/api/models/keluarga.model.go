package models

import (
	"errors"
	"math/rand"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateKeluarga(c echo.Context, k *entity.Keluarga) (entity.Keluarga, error) {
	db := db.GetDB(c)

	err := db.Create(&k)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Keluarga{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.Keluarga{}, errors.New("gagal membuat keluarga")
	}

	return *k, nil
}

func GetAllKeluarga(c echo.Context, filter string) ([]entity.Keluarga, error) {
	var keluargas []entity.Keluarga
	db := db.GetDB(c)
	var err *gorm.DB
	if filter != "" {
		err = db.Where("nama = ?", filter).Find(&keluargas)
	} else {
		err = db.Find(&keluargas)
	}
	if err.Error != nil {
		c.Logger().Error(err)
		return keluargas, err.Error
	}

	return keluargas, nil
}

func GetAllKeluargaWithEntity(c echo.Context, filter string, entitas string) ([]entity.Keluarga, error) {
	var keluargas []entity.Keluarga
	db := db.GetDB(c)
	var err *gorm.DB
	if filter != "" {
		err = db.Preload(entitas).Where("nama = ?", filter).Find(&keluargas)
	} else {
		err = db.Preload(entitas).Find(&keluargas)
	}
	if err.Error != nil {
		c.Logger().Error(err)
		return keluargas, err.Error
	}

	return keluargas, nil
}

func GetKeluargaByID(c echo.Context, id string) (entity.Keluarga, error) {
	var k entity.Keluarga
	db := db.GetDB(c)

	err := db.First(&k, "id = ?", id)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Keluarga{}, errors.New("id tidak ditemukan atau tidak valid")
	}

	return k, nil
}

func UpdateKeluargaById(c echo.Context, id string, k *entity.Keluarga) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.Keluarga{}).Where("id = ?", id).Updates(k)
	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeleteKeluargaById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Keluarga{})
	if err.Error != nil || err.RowsAffected == 0 {
		c.Logger().Error(err)
		return 0, err.Error
	}

	return err.RowsAffected, nil
}


func GetKeluargaByKode(c echo.Context, kode string) (entity.Keluarga, error) {
	var k entity.Keluarga
	db := db.GetDB(c)

	err := db.First(&k, "kode_keluarga = ?", kode)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Keluarga{}, errors.New("kode tidak ditemukan atau tidak valid")
	}

	return k, nil
}

func GenerateKodeKeluarga(c echo.Context, n int16) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	kode := string(b)
	db := db.GetDB(c)

	err := db.Where("kode_keluarga = ?", kode).First(&entity.Keluarga{})
	for err.Error == nil {
		b = make([]byte, n)
		for i := range b {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		}
		kode = string(b)
		err = db.Where("kode_keluarga = ?", kode).First(&entity.Keluarga{})
	}
	return kode
}
