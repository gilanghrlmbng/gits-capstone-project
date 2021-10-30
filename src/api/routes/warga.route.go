package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
)

func Warga(e *echo.Echo) *echo.Echo {

	e.POST("/warga", controllers.CreateWarga)
	e.GET("/warga", controllers.GetAllWarga)
	e.GET("/warga/:id", controllers.GetWargaByID)
	e.PUT("/warga/:id", controllers.UpdateWargaById)
	e.DELETE("/warga/:id", controllers.SoftDeleteWargaById)

	return e
}
