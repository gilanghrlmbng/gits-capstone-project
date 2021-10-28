package db

import (
	"fmt"
	"src/config"
	"src/migrations"
	"src/utils/errlogger"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB = nil
var err error

func Init(tableDelete, dataInitialization bool) {

	log.Info().Msg("menginisialisasikan database")

	config := config.GetConfig()
	log.Info().Msg(fmt.Sprintf("%v", config))

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.Username,
		config.Database.Password,
		config.Database.Name)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true, Logger: logger.Default.LogMode(logger.Info)})
	errlogger.ErrFatalPanic(err)

	if tableDelete {
		log.Info().Msg("menghapus tabel yang ada")
		db.Exec(`
DO $$ DECLARE
    r RECORD;
BEGIN
    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
        EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
    END LOOP;
END $$;`)
	}

	migrations.Migration(db)

	if dataInitialization {
		initData(db)
		// fmt.Print(" ")
	}

	log.Info().Msg("database terinisialisasi")
}

func GetDB() *gorm.DB {
	if db == nil {
		errlogger.FatalPanicMessage("db belum terinisilisasi")
	}
	return db
}

func initData(db *gorm.DB) {
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
	_ = SeedKeluarga(db, listIdRt)

	log.Info().Msg("dummy data terinisialisasi")
}
