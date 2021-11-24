package models

import (
	"errors"
	"src/db"
	"src/entity"

	"github.com/labstack/echo/v4"
)

func CreateOrder(c echo.Context, ord *entity.Order) (entity.Order, error) {
	db := db.GetDB(c)

	err := db.Create(&ord)

	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Order{}, err.Error
	}
	if err.RowsAffected == 0 {
		return entity.Order{}, errors.New("gagal menambahkan list order")
	}

	return *ord, nil
}

func GetAllOrder(c echo.Context) ([]entity.Order, error) {
	var ord []entity.Order
	db := db.GetDB(c)

	err := db.Find(&ord)
	if err.Error != nil {
		c.Logger().Error(err)
		return ord, err.Error
	}
	return ord, nil
}

func GetOrderByID(c echo.Context, id string) (entity.Order, error) {
	var ord entity.Order
	db := db.GetDB(c)

	err := db.First(&ord, "id = ?", id)
	if err.Error != nil {
		c.Logger().Error(err)
		return entity.Order{}, errors.New("id tidak ditemukan")
	}
	return ord, nil
}

func UpdateOrderById(c echo.Context, id string, order *entity.Order) (int64, error) {
	db := db.GetDB(c)

	err := db.Model(&entity.Order{}).Where("id = ?", id).Updates(order)

	if err.Error != nil {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}

func SoftDeleteOrderById(c echo.Context, id string) (int64, error) {
	db := db.GetDB(c)

	err := db.Where("id = ?", id).Delete(&entity.Order{})

	if err.Error != nil || err.RowsAffected == 0 {
		c.Logger().Error(err)
		return 0, err.Error
	}
	return err.RowsAffected, nil
}
