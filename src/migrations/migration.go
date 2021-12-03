package migrations

import (
	"src/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Migration(e *echo.Echo, db *gorm.DB) {
	/*
		Please fill the params in AutoMigrate with your entity
		so you will see
		db.AutoMigrate(&Entity1{}, &Entity2{}, &Entity3{}, ...)
	*/

	e.Logger.Info("Memulai dengan automigrate")

	err := db.AutoMigrate(&entity.Rt{}, &entity.PengurusRT{}, &entity.Keluarga{}, &entity.Warga{}, &entity.Tagihan{}, &entity.Produk{}, &entity.Order{}, &entity.ItemOrder{}, &entity.Pembayaran{}, &entity.Informasi{}, &entity.DompetRT{}, &entity.Persuratan{}, &entity.DompetKeluarga{})

	if err != nil {
		e.Logger.Error(err)

	}
}

func DeleteAllTable(e *echo.Echo, db *gorm.DB) {
	e.Logger.Info("Mereset Semua Tabel")
	err := db.Migrator().DropTable(&entity.Rt{}, &entity.PengurusRT{}, &entity.Keluarga{}, &entity.Warga{}, &entity.Tagihan{}, &entity.Produk{}, &entity.Order{}, &entity.ItemOrder{}, &entity.Pembayaran{}, &entity.Informasi{}, &entity.DompetRT{}, &entity.Persuratan{}, &entity.DompetKeluarga{})
	if err != nil {
		e.Logger.Error(err)

	}
}
