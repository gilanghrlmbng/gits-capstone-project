package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func DompetRT(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {
	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(JWTconfig))
	auth.POST("/dompet", controllers.CreateDompet)
	auth.GET("/dompet", controllers.GetAllDompet)
	auth.GET("/dompet/:id", controllers.GetDompetByID)
	auth.PUT("/dompet/:id", controllers.UpdateDompetById)
	auth.DELETE("/dompet/:id", controllers.SoftDeleteDompetById)

	return e
}
