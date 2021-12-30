package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
)

func CreatePembayaran(c echo.Context, pem *entity.Pembayaran) (entity.Pembayaran, error) {
	db := db.GetDB(c)

	err := db.Create(&pem)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Pembayaran{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.Pembayaran{}, errors.New("gagal mencatat pembayaran")
	}

	return *pem, nil
}

func UpdatePembayaranById(c echo.Context, id string, pem *entity.Pembayaran) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.Pembayaran{}).Where("id = ?", id).Updates(pem)

	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
