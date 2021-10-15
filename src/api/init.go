package api

import (
	"src/api/routes"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func Init() *echo.Echo {
	log.Info().Msg("menginisialisasikan server")

	e := echo.New()
	e = routes.Init(e)

	log.Info().Msg("server terinisialisasi")

	return e
}
