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
		db.Exec(`DROP SCHEMA public CASCADE;
		CREATE SCHEMA public; GRANT ALL ON SCHEMA public TO postgres;
		GRANT ALL ON SCHEMA public TO public;`)
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

	// // Tipe User
	// data, err = os.ReadFile("db/dummy/tipe_user.sql")
	// errlogger.ErrFatalPanic(err)
	// db.Exec(string(data))

	// // User
	// data, err = os.ReadFile("db/dummy/user.sql")
	// errlogger.ErrFatalPanic(err)
	// db.Exec(string(data))

	// // Post
	// data, err = os.ReadFile("db/dummy/post.sql")
	// errlogger.ErrFatalPanic(err)
	// db.Exec(string(data))

	// // Komentar
	// data, err = os.ReadFile("db/dummy/komentar.sql")
	// errlogger.ErrFatalPanic(err)
	// db.Exec(string(data))

	log.Info().Msg("dummy data terinisialisasi")
}
