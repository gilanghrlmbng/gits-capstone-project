package models

import (
	"errors"
	"src/db"
	"src/entity"

	"gorm.io/gorm"
)

func CreateKeluarga(k *entity.Keluarga) (entity.Keluarga, error) {
	db := db.GetDB()

	err := db.Create(&k)
	if err.Error != nil {
		return entity.Keluarga{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.Keluarga{}, errors.New("gagal membuat keluarga")
	}

	return *k, nil
}

func GetAllKeluarga(filter string) ([]entity.Keluarga, error) {
	var keluargas []entity.Keluarga
	db := db.GetDB()
	var err *gorm.DB
	if filter != "" {
		err = db.Where("nama = ?", filter).Find(&keluargas)
	} else {
		err = db.Find(&keluargas)
	}
	if err.Error != nil {
		return keluargas, err.Error
	}

	return keluargas, nil
}

func GetKeluargaByID(id string) (entity.Keluarga, error) {
	var k entity.Keluarga
	db := db.GetDB()

	err := db.First(&k, "id = ?", id)
	if err.Error != nil {
		return entity.Keluarga{}, errors.New("id tidak ditemukan atau tidak valid")
	}

	return k, nil
}

func UpdateKeluargaById(id string, k *entity.Keluarga) (int64, error) {
	db := db.GetDB()

	err := db.Model(&entity.Keluarga{}).Where("id = ?", id).Updates(k)
	if err.Error != nil {
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeleteKeluargaById(id string) (int64, error) {
	db := db.GetDB()

	err := db.Where("id = ?", id).Delete(&entity.Keluarga{})
	if err.Error != nil || err.RowsAffected == 0 {
		return 0, err.Error
	}

	return err.RowsAffected, nil
}
