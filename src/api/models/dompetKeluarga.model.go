package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateDompetKeluarga(c echo.Context, d *entity.DompetKeluarga) (entity.DompetKeluarga, error) {
	db := db.GetDB(c)

	err := db.Create(&d)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.DompetKeluarga{}, err.Error
	}

	if err.RowsAffected == 0 {
		return entity.DompetKeluarga{}, errors.New("gagal menambahkan dompet")
	}

	return *d, nil
}

func GetAllDompetKeluarga(c echo.Context) ([]entity.DompetKeluarga, error) {
	var dompet []entity.DompetKeluarga
	db := db.GetDB(c)

	err := db.Find(&dompet)
	if err.Error != nil {
		c.Logger().Error(err)
		return dompet, err.Error
	}

	return dompet, nil
}

func GetDompetKeluargaByID(c echo.Context, id, id_keluarga string) (entity.DompetKeluarga, error) {
	var d entity.DompetKeluarga
	var err *gorm.DB
	db := db.GetDB(c)

	if id_keluarga != "" {
		err = db.First(&d, "id_keluarga = ?", id_keluarga)
	} else {
		err = db.First(&d, "id = ?", id)
	}

	if err.Error != nil {
		c.Logger().Error(err)
		return entity.DompetKeluarga{}, errors.New("id tidak ditemukan atau tidak valid")
	}
	return d, nil
}

func UpdateDompetKeluargaById(c echo.Context, id string, dompet *entity.DompetKeluarga) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.DompetKeluarga{}).Where("id = ?", id).Updates(dompet)

	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeleteDompetKeluargaById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.DompetKeluarga{})
	if err.Error != nil || err.RowsAffected == 0 {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
