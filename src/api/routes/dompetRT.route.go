package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func DompetRT(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {
	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(JWTconfig))
	auth.GET("/dompetrt", controllers.GetAllDompet)
	auth.GET("/dompetrt/me", controllers.GetDompetByID)
	auth.GET("/dompetrt/:id", controllers.GetDompetByID)
	auth.PUT("/dompetrt/:id", controllers.UpdateDompetById)
	auth.DELETE("/dompetrt/:id", controllers.SoftDeleteDompetById)

	return e
}
