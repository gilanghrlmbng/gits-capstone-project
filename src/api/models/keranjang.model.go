package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateKeranjang(c echo.Context, ord *entity.Keranjang) (entity.Keranjang, error) {
	db := db.GetDB(c)

	err := db.Create(&ord)

	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Keranjang{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.Keranjang{}, errors.New("gagal menambahkan list order")
	}

	return *ord, nil
}

func GetAllKeranjang(c echo.Context, idPembeli, idPenjual string) ([]entity.Keranjang, error) {
	var ord []entity.Keranjang
	db := db.GetDB(c)
	var err *gorm.DB
	if idPembeli != "" {
		err = db.Preload("ItemKeranjang").Order("id desc").Where("id_warga = ?", idPembeli).Find(&ord)
	} else if idPenjual != "" {
		err = db.Preload("ItemKeranjang").Order("id desc").Where("id_keluarga_penjual = ?", idPenjual).Find(&ord)
	} else {
		err = db.Preload("ItemKeranjang").Order("id desc").Find(&ord)
	}
	if err.Error != nil {
		c.Logger().Error(err)
		return ord, err.Error
	}
	return ord, nil
}

func GetKeranjangByID(c echo.Context, id string) (entity.Keranjang, error) {
	var ord entity.Keranjang
	db := db.GetDB(c)

	err := db.Preload("ItemKeranjang").First(&ord, "id = ?", id)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Keranjang{}, errors.New("id tidak ditemukan")
	}
	return ord, nil
}

func GetKeranjangByIDWarga(c echo.Context, id_warga string) (entity.Keranjang, error) {
	var ord entity.Keranjang
	db := db.GetDB(c)

	err := db.Preload("ItemKeranjang").Order("id desc").First(&ord, "id_warga = ?", id_warga)
	if err.Error != nil {
		c.Logger().Error(err)
		return ord, errors.New("id tidak ditemukan")
	}
	return ord, nil
}

func UpdateKeranjangById(c echo.Context, id string, order *entity.Keranjang) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.Keranjang{}).Where("id = ?", id).Updates(order)

	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeleteKeranjangById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Keranjang{})

	if err.Error != nil || err.RowsAffected == 0 {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
