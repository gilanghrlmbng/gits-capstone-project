package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
)

func PengurusRT(e *echo.Echo) *echo.Echo {
	e.POST("/pengurus", controllers.CreatePengurus)
	e.GET("/pengurus", controllers.GetAllPengurusRT)
	e.GET("/pengurus/:id", controllers.GetPengurusByID)
	e.PUT("/pengurus/:id", controllers.UpdatePengurusById)
	e.DELETE("/pengurus/:id", controllers.SoftDeletePengurusById)

	return e
}
