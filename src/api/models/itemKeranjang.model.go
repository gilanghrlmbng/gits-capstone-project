package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
)

func CreateItemKeranjang(c echo.Context, i *entity.ItemKeranjang) (entity.ItemKeranjang, error) {
	db := db.GetDB(c)

	err := db.Create(&i)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.ItemKeranjang{}, err.Error
	}

	if err.RowsAffected == 0 {
		return entity.ItemKeranjang{}, errors.New("gagal menambahkan item order")
	}

	return *i, nil
}

func CreateBatchItemKeranjang(c echo.Context, items []entity.ItemKeranjang) ([]entity.ItemKeranjang, error) {
	db := db.GetDB(c)

	err := db.Create(items)
	if err.Error != nil {
		c.Logger().Error(err)
		return []entity.ItemKeranjang{}, err.Error
	}

	if err.RowsAffected == 0 {
		return []entity.ItemKeranjang{}, errors.New("gagal menambahkan item order")
	}

	return items, nil
}

func GetAllItemKeranjang(c echo.Context) ([]entity.ItemKeranjang, error) {
	var item []entity.ItemKeranjang
	db := db.GetDB(c)

	err := db.Find(&item)
	if err.Error != nil {
		c.Logger().Error(err)
		return item, err.Error
	}
	return item, nil
}

func GetItemKeranjangById(c echo.Context, id string) ([]entity.ItemKeranjang, error) {
	var item []entity.ItemKeranjang
	db := db.GetDB(c)

	err := db.First(&item, "id = ?", id)
	if err.Error != nil {
		c.Logger().Error(err)
		return item, errors.New("id tidak ditemukan atau tidak valid")
	}
	return item, nil
}

func GetItemKeranjangByIDKeranjang(c echo.Context, id_order string) ([]entity.ItemKeranjang, error) {
	var item []entity.ItemKeranjang
	db := db.GetDB(c)

	err := db.First(&item, "id_keranjang = ?", id_order)
	if err.Error != nil {
		c.Logger().Error(err)
		return item, errors.New("id tidak ditemukan")
	}
	return item, nil
}

func HardDeleteItemKeranjang(c echo.Context, id_order string) error {
	item := entity.ItemKeranjang{
		Id: id_order,
	}
	db := db.GetDB(c)

	err := db.Unscoped().Delete(&item)
	if err.Error != nil {
		c.Logger().Error(err)
		return err.Error
	}
	return nil
}
