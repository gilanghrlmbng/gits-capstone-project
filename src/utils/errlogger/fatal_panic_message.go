package errlogger

import "github.com/rs/zerolog/log"

func FatalPanicMessage(msg string) {
	log.Fatal().Msg(msg)
}
