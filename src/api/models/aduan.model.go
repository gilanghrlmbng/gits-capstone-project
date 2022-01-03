package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateAduan(c echo.Context, a *entity.Aduan) (entity.Aduan, error) {
	db := db.GetDB(c)

	err := db.Create(&a)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Aduan{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.Aduan{}, errors.New("gagal menambahkan Aduan")
	}

	return *a, nil
}

func GetAllAduan(c echo.Context, idWarga, IdRT string) ([]entity.Aduan, error) {
	var aduans []entity.Aduan
	db := db.GetDB(c)

	var err *gorm.DB
	if idWarga != "" {
		err = db.Where("id_warga = ?", idWarga).Find(&aduans)
	} else if IdRT != "" {
		err = db.Where("id_rt = ?", IdRT).Find(&aduans)
	} else {
		err = db.Find(&aduans)
	}

	if err.Error != nil {
		c.Logger().Error(err)
		return aduans, err.Error
	}

	return aduans, nil
}

func GetAduanByID(c echo.Context, id string) (entity.Aduan, error) {
	var a entity.Aduan
	db := db.GetDB(c)

	err := db.First(&a, "id = ? ", id)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Aduan{}, errors.New("id tidak ditemukan atau tidak valid")
	}
	return a, nil
}

func UpdateAduanById(c echo.Context, id string, a *entity.Aduan) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.Aduan{}).Where("id = ?", id).Updates(a)
	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeleteAduanById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Aduan{})
	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
