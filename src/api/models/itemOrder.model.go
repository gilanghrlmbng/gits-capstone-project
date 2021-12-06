package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
)

func CreateItemOrder(c echo.Context, i *entity.ItemOrder) (entity.ItemOrder, error) {
	db := db.GetDB(c)

	err := db.Create(&i)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.ItemOrder{}, err.Error
	}

	if err.RowsAffected == 0 {
		return entity.ItemOrder{}, errors.New("gagal menambahkan item order")
	}

	return *i, nil
}

func GetAllItemOrder(c echo.Context) ([]entity.ItemOrder, error) {
	var item []entity.ItemOrder
	db := db.GetDB(c)

	err := db.Find(&item)
	if err.Error != nil {
		c.Logger().Error(err)
		return item, err.Error
	}
	return item, nil
}

func GetItemOrderByID(c echo.Context, id string) ([]entity.ItemOrder, error) {
	var item []entity.ItemOrder
	db := db.GetDB(c)

	err := db.First(&item, "id = ?", id)
	if err.Error != nil {
		c.Logger().Error(err)
		return item, errors.New("id tidak ditemukan atau tidak valid")
	}
	return item, nil
}

func ProdukSearch(c echo.Context, nama string) (entity.Produk, error) {

	var prt entity.Produk
	db := db.GetDB(c)

	err := db.First(&prt, "nama = ?", nama)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Produk{}, errors.New("nama tidak ditemukan")
	}
	return prt, nil
}
