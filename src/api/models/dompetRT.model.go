package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateDompet(c echo.Context, d *entity.DompetRT) (entity.DompetRT, error) {
	db := db.GetDB(c)

	err := db.Create(&d)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.DompetRT{}, err.Error
	}

	if err.RowsAffected == 0 {
		return entity.DompetRT{}, errors.New("gagal menambahkan dompet")
	}

	return *d, nil
}

func GetAllDompet(c echo.Context) ([]entity.DompetRT, error) {
	var dompet []entity.DompetRT
	db := db.GetDB(c)

	err := db.Find(&dompet)
	if err.Error != nil {
		c.Logger().Error(err)
		return dompet, err.Error
	}

	return dompet, nil
}

func GetDompetByID(c echo.Context, id, id_rt string) (entity.DompetRT, error) {
	var d entity.DompetRT
	var err *gorm.DB
	db := db.GetDB(c)
	if id_rt != "" {
		err = db.First(&d, "id_rt = ?", id_rt)
	} else {
		err = db.First(&d, "id = ?", id)
	}
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.DompetRT{}, errors.New("id tidak ditemukan atau tidak valid")
	}
	return d, nil
}

func UpdateDompetById(c echo.Context, id string, dompet *entity.DompetRT) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.DompetRT{}).Where("id = ?", id).Updates(dompet)

	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeleteDompetById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.DompetRT{})
	if err.Error != nil || err.RowsAffected == 0 {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
