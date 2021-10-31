package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
)

func Keluarga(e *echo.Echo) *echo.Echo {

	e.POST("/keluarga", controllers.CreateKeluarga)
	e.GET("/keluarga", controllers.GetAllKeluarga)
	e.GET("/keluarga/warga", controllers.GetAllKeluargaWithWarga)
	e.GET("/keluarga/:id", controllers.GetKeluargaByID)
	e.PUT("/keluarga/:id", controllers.UpdateKeluargaById)
	e.DELETE("/keluarga/:id", controllers.SoftDeleteKeluargaById)

	return e
}
