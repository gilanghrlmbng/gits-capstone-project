package main

import (
	"os"
	"src/api"
	"src/db"
	"src/utils/errlogger"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Inisialisasi Logger
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Inisialisasi Env
	err := godotenv.Load()
	errlogger.ErrFatalPanic(err)

	// Inisialisasi DB
	db.Init(true)
	// Inisialisasi Server
	e := api.Init()

	// Server Listener
	e.Logger.Fatal((e.Start(":4132")))
}
