package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateProduk(c echo.Context, p *entity.Produk) (entity.Produk, error) {
	db := db.GetDB(c)

	err := db.Create(&p)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Produk{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.Produk{}, errors.New("gagal menambahkan produk")
	}

	return *p, nil
}

func GetAllProduk(c echo.Context, idKeluarga string) (p []entity.Produk, err error) {
	var produks []entity.Produk
	db := db.GetDB(c)
	var errs *gorm.DB
	if idKeluarga != "" {
		errs = db.Where("id_keluarga = ?", idKeluarga).Find(&produks)
	} else {
		errs = db.Find(&produks)
	}

	if errs.Error != nil {
		c.Logger().Error(err)
		err = errs.Error
		return
	}

	return produks, nil
}

func GetProdukByID(c echo.Context, id string) (entity.Produk, error) {
	var p entity.Produk
	db := db.GetDB(c)

	err := db.First(&p, "id = ?", id)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Produk{}, errors.New("id tidak ditemukan atau tidak valid")
	}
	return p, nil
}

func UpdateProdukById(c echo.Context, id string, produk *entity.Produk) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.Produk{}).Where("id = ? ", id).Updates(produk)

	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeleteProdukById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Produk{})

	if err.Error != nil || err.RowsAffected == 0 {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
