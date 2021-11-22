package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
)

func DompetRT(e *echo.Echo) *echo.Echo {
	e.POST("/dompet", controllers.CreateDompet)
	e.GET("/dompet", controllers.GetAllDompet)
	e.GET("/dompet/:id", controllers.GetDompetByID)
	e.PUT("/dompet/:id", controllers.UpdateDompetById)
	e.DELETE("/dompet/:id", controllers.SoftDeleteDompetById)

	return e
}
