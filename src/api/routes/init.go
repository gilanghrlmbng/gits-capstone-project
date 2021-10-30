package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) *echo.Echo {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
	e.Logger.Info("menginisialisasikan routes")
	e = Keluarga(e)
	e = RT(e)
	e = PengurusRT(e)
	e = Warga(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Worlds!!!")
	})

	e.Logger.Info("routes terinisialisasi")

	return e
}
