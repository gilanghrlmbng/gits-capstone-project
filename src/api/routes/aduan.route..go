package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Aduan(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {
	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(JWTconfig))
	auth.POST("/aduan", controllers.CreateAduan)
	auth.GET("/aduan", controllers.GetAllAduan)
	auth.GET("/aduan/:id", controllers.GetAduanByID)
	auth.PUT("/aduan/:id", controllers.UpdateAduanById)
	auth.DELETE("/aduan/:id", controllers.SoftDeleteAduanById)

	return e
}
