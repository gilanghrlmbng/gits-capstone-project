package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Warga(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {
	auth := e.Group("/warga")
	auth.Use(middleware.JWTWithConfig(JWTconfig))
	
	auth.GET("", controllers.GetAllWarga)
	auth.GET("/me", controllers.GetWargaByID)
	auth.GET("/detail/:id", controllers.GetWargaByID)
	auth.PUT("/:id", controllers.UpdateWargaById)
	auth.DELETE("/:id", controllers.SoftDeleteWargaById)

	e.POST("/warga", controllers.CreateWarga)
	e.POST("/warga/login", controllers.LoginWarga)

	return e
}
