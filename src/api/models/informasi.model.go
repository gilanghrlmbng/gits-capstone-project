package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateInformasi(c echo.Context, p *entity.Informasi) (entity.Informasi, error) {
	db := db.GetDB(c)

	err := db.Create(&p)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Informasi{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.Informasi{}, errors.New("gagal menambahkan informasi")
	}

	return *p, nil
}

func GetAllInformasi(c echo.Context, idInformasi string) (p []entity.Informasi, err error) {
	var informasis []entity.Informasi
	db := db.GetDB(c)
	var errs *gorm.DB
	if idInformasi != "" {
		errs = db.Where("id_informasi = ?", idInformasi).Find(&informasis) // ini
	} else {
		errs = db.Find(&informasis)
	}

	if errs.Error != nil {
		c.Logger().Error(err)
		err = errs.Error
		return
	}

	return informasis, nil
}

func GetInformasiByID(c echo.Context, id string) (entity.Informasi, error) {
	var p entity.Informasi
	db := db.GetDB(c)

	err := db.First(&p, "id = ?", id)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Informasi{}, errors.New("id tidak ditemukan atau tidak valid")
	}
	return p, nil
}

func GetAllInformasiByKategori(c echo.Context, id_rt, kategori string) ([]entity.Informasi, error) {
	var informasi []entity.Informasi
	var err *gorm.DB
	db := db.GetDB(c)
	err = db.Where("id_rt = ? AND kategori = ?", id_rt, kategori).Find(&informasi)
	if err.Error != nil {
		c.Logger().Error(err)
		return informasi, err.Error
	}
	return informasi, nil
}

func UpdateInformasiById(c echo.Context, id string, informasi *entity.Informasi) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.Informasi{}).Where("id = ? ", id).Updates(informasi) // ini

	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeleteInformasiById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Informasi{})

	if err.Error != nil || err.RowsAffected == 0 {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
