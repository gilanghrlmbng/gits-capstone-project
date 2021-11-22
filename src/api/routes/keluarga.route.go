package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Keluarga(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {

	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(JWTconfig))
	auth.POST("/keluarga", controllers.CreateKeluarga)
	auth.GET("/keluargasaya", controllers.GetKeluargaByID)
	auth.GET("/keluarga", controllers.GetAllKeluarga)
	auth.GET("/keluarga/warga", controllers.GetAllKeluargaWithWarga)
	auth.PUT("/keluarga/:id", controllers.UpdateKeluargaById)
	auth.GET("/keluarga/:id", controllers.GetKeluargaByID)
	auth.DELETE("/keluarga/:id", controllers.SoftDeleteKeluargaById)

	return e
}
