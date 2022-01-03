package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateTagihan(c echo.Context, s *entity.Tagihan) (entity.Tagihan, error) {
	db := db.GetDB(c)

	err := db.Create(&s)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Tagihan{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.Tagihan{}, errors.New("gagal menambahkan Tagihan")
	}

	return *s, nil
}

func GetAllTagihan(c echo.Context, idKeluarga, idRT, terbayar string) (s []entity.Tagihan, err error) {
	var tagihan []entity.Tagihan
	db := db.GetDB(c)
	var errs *gorm.DB

	if idKeluarga != "" && terbayar != "" {
		errs = db.Where("id_keluarga = ? AND terbayar = ?", idKeluarga, terbayar).Find(&tagihan)
	} else if idKeluarga != "" {
		errs = db.Where("id_keluarga = ?", idKeluarga).Find(&tagihan)
	} else if idRT != "" && terbayar != "" {
		errs = db.Where("id_rt = ? AND terbayar = ?", idRT, terbayar).Find(&tagihan)
	} else if idRT != "" {
		errs = db.Where("id_rt = ?", idRT).Find(&tagihan)
	} else {
		errs = db.Find(&tagihan)
	}

	if errs.Error != nil {
		c.Logger().Error(err)
		err = errs.Error
		return
	}

	return tagihan, nil
}

func GetTagihanByID(c echo.Context, id string) (entity.Tagihan, error) {
	var s entity.Tagihan
	db := db.GetDB(c)

	err := db.First(&s, "id = ?", id)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Tagihan{}, errors.New("id tidak ditemukan atau tidak valid")
	}
	return s, nil
}

func UpdateTagihanById(c echo.Context, id string, tagihan *entity.Tagihan) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.Tagihan{}).Where("id = ? ", id).Updates(tagihan)

	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeleteTagihanById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Tagihan{})

	if err.Error != nil || err.RowsAffected == 0 {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
