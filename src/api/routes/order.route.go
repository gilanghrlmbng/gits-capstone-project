package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Order(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {
	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(JWTconfig))
	auth.POST("/order", controllers.CreateOrder)
	auth.GET("/order", controllers.GetAllOrder)
	auth.GET("/order/:id", controllers.GetOrderByID)
	auth.PUT("/order/:id", controllers.UpdateOrderById)
	auth.DELETE("/order/:id", controllers.SoftDeleteOrderById)

	return e
}
