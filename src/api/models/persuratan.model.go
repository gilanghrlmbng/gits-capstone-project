package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreatePersuratan(c echo.Context, s *entity.Persuratan) (entity.Persuratan, error) {
	db := db.GetDB(c)

	err := db.Create(&s)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Persuratan{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.Persuratan{}, errors.New("gagal menambahkan Persuratan")
	}

	return *s, nil
}

func GetAllPersuratan(c echo.Context, idRT string) (s []entity.Persuratan, err error) {
	var persuratans []entity.Persuratan
	db := db.GetDB(c)
	var errs *gorm.DB
	if idRT == "" {
		errs = db.Find(&persuratans)
	} else {
		errs = db.Where("id_rt = ?", idRT).Find(&persuratans)
	}

	if errs.Error != nil {
		c.Logger().Error(err)
		err = errs.Error
		return
	}

	return persuratans, nil
}

func GetPersuratanByID(c echo.Context, id string) (entity.Persuratan, error) {
	var s entity.Persuratan
	db := db.GetDB(c)

	err := db.First(&s, "id = ?", id)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Persuratan{}, errors.New("id tidak ditemukan atau tidak valid")
	}
	return s, nil
}

func UpdatePersuratanById(c echo.Context, id string, persuratan *entity.Persuratan) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.Persuratan{}).Where("id = ? ", id).Updates(persuratan)

	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeletePersuratanById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Persuratan{})

	if err.Error != nil || err.RowsAffected == 0 {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
