package migrations

import (
	"src/entity"
	"src/utils/errlogger"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	/*
		Please fill the params in AutoMigrate with your entity
		so you will see
		db.AutoMigrate(&Entity1{}, &Entity2{}, &Entity3{}, ...)
	*/

	log.Info().Msg("memulai dengan automigrate")

	err := db.AutoMigrate(&entity.Rt{}, &entity.PengurusRT{}, &entity.Keluarga{}, &entity.Warga{}, &entity.Tagihan{}, &entity.Produk{})

	errlogger.ErrFatalPanic(err)
}
