package api

import (
	"src/api/routes"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) *echo.Echo {
	e.Logger.Info("menginisialisasikan server")

	e = routes.Init(e)

	e.Logger.Info("server terinisialisasi")

	return e
}
