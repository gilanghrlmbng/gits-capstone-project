package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func DompetKeluarga(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {
	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(JWTconfig))
	auth.GET("/dompetkeluarga", controllers.GetAllDompetKeluarga)
	auth.GET("/dompetkeluarga/me", controllers.GetDompetKeluargaByID)
	auth.GET("/dompetkeluarga/:id", controllers.GetDompetKeluargaByID)
	auth.PUT("/dompetkeluarga/:id", controllers.UpdateDompetKeluargaById)
	auth.DELETE("/dompetkeluarga/:id", controllers.SoftDeleteDompetKeluargaById)

	return e
}
