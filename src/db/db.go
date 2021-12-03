package db

import (
	"fmt"
	"src/config"
	"src/migrations"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil
var err error

func Init(e *echo.Echo, tableDelete, dataInitialization bool) {

	e.Logger.Info("menginisialisasikan database")

	config := config.GetConfig(e)
	e.Logger.Info(config)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		// DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		e.Logger.Error(err)

	}

	if tableDelete {
		migrations.DeleteAllTable(e, db)
	}

	migrations.Migration(e, db)

	if dataInitialization {
		initData(e, db)
		// fmt.Print(" ")
	}

	e.Logger.Info("database terinisialisasi")
}

func GetDB(c echo.Context) *gorm.DB {
	if db == nil {
		c.Logger().Error("db belum terinisilisasi")
	}
	return db
}

func initData(e *echo.Echo, db *gorm.DB) {
	/*
		Use this function to make a initial data.
		You need to initialize your data first and the loop through the data.
		To Create Record please refer reading this https://gorm.io/docs/create.html
	*/

	// RT
	listIdRt := SeedRT(db)

	// Pengurus RT
	_ = SeedPengurusRT(db, listIdRt)

	// Keluarga
	listIdKeluarga := SeedKeluarga(db, listIdRt)

	//Warga
	_ = SeedWarga(db, listIdKeluarga)

	// Produk
	_ = SeedProduk(db, listIdKeluarga)

	// DompetRT
	_ = SeedDompetRT(db, listIdRt)

	// Dompet Keluarga
	_ = SeedDompetKeluarga(db, listIdKeluarga)

	e.Logger.Info("dummy data terinisialisasi")
}
