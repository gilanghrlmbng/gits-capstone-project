package errlogger

import "github.com/rs/zerolog/log"

func ErrFatalPanic(err error) {
	if err != nil {
		log.Fatal().Err(err)
		panic(err)
	}
}
