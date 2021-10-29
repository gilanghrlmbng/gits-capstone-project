package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func Init(e *echo.Echo) *echo.Echo {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
	log.Info().Msg("menginisialisasikan routes")
	e = Keluarga(e)
	e = RT(e)
	e = PengurusRT(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Worlds!!!")
	})

	log.Info().Msg("routes terinisialisasi")

	return e
}
