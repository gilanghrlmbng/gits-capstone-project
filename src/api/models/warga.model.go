package models

import (
	"errors"
	"fmt"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateWarga(c echo.Context, w *entity.Warga) (entity.Warga, error) {
	db := db.GetDB(c)
	if w.Gambar == "" {
		w.Gambar = fmt.Sprintf("https://dummyimage.com/500x500/eee/fff&text=%c", w.Nama[0])
	}
	err := db.Create(&w)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Warga{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.Warga{}, errors.New("gagal membuat keluarga")
	}

	return *w, nil
}

func GetAllWarga(c echo.Context, IdKeluarga string) ([]entity.Warga, error) {
	var ws []entity.Warga
	db := db.GetDB(c)

	var err *gorm.DB
	if IdKeluarga != "" {
		err = db.Where("id_keluarga = ?", IdKeluarga).Find(&ws)
	} else {
		err = db.Find(&ws)
	}

	if err.Error != nil {
		c.Logger().Error(err)
		return ws, err.Error
	}

	return ws, nil
}

func GetWargaByID(c echo.Context, id string) (entity.Warga, error) {
	var w entity.Warga
	db := db.GetDB(c)

	err := db.First(&w, "id = ?", id)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Warga{}, errors.New("id tidak ditemukan atau tidak valid")
	}

	return w, nil
}

func GetWargaByEmail(c echo.Context, email string) (entity.Keluarga, error) {
	// var w entity.Warga
	var k entity.Keluarga
	db := db.GetDB(c)
	err := db.Preload("Warga", "email = ?", email).First(&k)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Keluarga{}, errors.New("email tidak ditemukan")
	}

	return k, nil
}

func UpdateWargaById(c echo.Context, id string, w *entity.Warga) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.Warga{}).Where("id = ?", id).Updates(w)
	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeleteWargaById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Warga{})
	if err.Error != nil || err.RowsAffected == 0 {
		c.Logger().Error(err)
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
