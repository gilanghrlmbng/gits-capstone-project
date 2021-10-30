package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
)

func CreateWarga(c echo.Context, w *entity.Warga) (entity.Warga, error) {
	db := db.GetDB(c)

	err := db.Create(&w)
	if err.Error != nil {
		return entity.Warga{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.Warga{}, errors.New("gagal membuat keluarga")
	}

	return *w, nil
}

func GetAllWarga(c echo.Context) ([]entity.Warga, error) {
	var ws []entity.Warga
	db := db.GetDB(c)

	err := db.Find(&ws)
	if err.Error != nil {
		return ws, err.Error
	}

	return ws, nil
}

func GetWargaByID(c echo.Context, id string) (entity.Warga, error) {
	var w entity.Warga
	db := db.GetDB(c)

	err := db.First(&w, "id = ?", id)
	if err.Error != nil {
		return entity.Warga{}, errors.New("id tidak ditemukan atau tidak valid")
	}

	return w, nil
}

func GetWargaByEmail(c echo.Context, email string) (entity.Warga, error) {
	var w entity.Warga
	db := db.GetDB(c)

	err := db.First(&w, "email = ?", email)
	if err.Error != nil {
		return entity.Warga{}, errors.New("email tidak ditemukan")
	}

	return w, nil
}

func UpdateWargaById(c echo.Context, id string, w *entity.Warga) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.Warga{}).Where("id = ?", id).Updates(w)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeleteWargaById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Warga{})
	if err.Error != nil || err.RowsAffected == 0 {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
