package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Tagihan(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {
	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(JWTconfig))
	auth.POST("/tagihan", controllers.CreateTagihan)
	auth.GET("/tagihan", controllers.GetAllTagihan)
	auth.GET("/tagihan/:id", controllers.GetTagihanByID)
	auth.PUT("/tagihan/:id", controllers.BayarTagihanByID)
	auth.DELETE("/tagihan/:id", controllers.SoftDeleteTagihanById)

	return e
}
