package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
)

func CreatePengurusRT(c echo.Context, prt *entity.PengurusRT) (entity.PengurusRT, error) {
	db := db.GetDB(c)

	err := db.Create(&prt)
	if err.Error != nil {
		return entity.PengurusRT{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.PengurusRT{}, errors.New("gagal membuat pengurus RT")
	}

	return *prt, nil
}

func GetAllPengurusRT(c echo.Context) ([]entity.PengurusRT, error) {
	var pengurusRTs []entity.PengurusRT
	db := db.GetDB(c)

	err := db.Find(&pengurusRTs)
	if err.Error != nil {
		return pengurusRTs, err.Error
	}

	return pengurusRTs, nil
}

func GetPengurusByID(c echo.Context, id string) (entity.PengurusRT, error) {
	var prt entity.PengurusRT
	db := db.GetDB(c)

	err := db.First(&prt, "id = ?", id)
	if err.Error != nil {
		return entity.PengurusRT{}, errors.New("id tidak ditemukan atau tidak valid")
	}
	return prt, nil
}

func UpdatePengurusById(c echo.Context, id string, pengurusRT *entity.PengurusRT) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.PengurusRT{}).Where("id = ? ", id).Updates(pengurusRT)

	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeletePengurusById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.PengurusRT{})

	if err.Error != nil || err.RowsAffected == 0 {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
