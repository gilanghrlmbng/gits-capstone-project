package routes

import (
	"net/http"
	"src/config"
	"src/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init(e *echo.Echo) *echo.Echo {
	JWTconfig := middleware.JWTConfig{
		TokenLookup: "header:Authorization",
		Claims:      &utils.JWTCustomClaims{},
		SigningKey:  []byte(config.GetConfig(e).Secret),
	}
	middleware.ErrJWTMissing.Code = http.StatusUnauthorized

	middleware.ErrJWTMissing.Message = utils.Error{
		Code:    http.StatusUnauthorized,
		Message: "Unauthorized",
	}
	middleware.ErrJWTInvalid.Message = utils.Error{
		Code:    http.StatusUnauthorized,
		Message: "Token Invalid",
	}
	e.Logger.Info("menginisialisasikan routes")
	e = Keluarga(e, JWTconfig)
	e = RT(e)
	e = PengurusRT(e)
	e = Warga(e, JWTconfig)
	e = Produk(e)
	e = DompetRT(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Worlds!!!")
	})

	e.Logger.Info("routes terinisialisasi")

	return e
}
