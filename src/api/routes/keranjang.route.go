package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Keranjang(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {
	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(JWTconfig))
	auth.PUT("/cart", controllers.UpdateKeranjangByIdWarga)
	auth.PUT("/cart/add", controllers.TambahItemKeranjang)
	auth.PUT("/cart/:id", controllers.UpdateKeranjangByIdWarga)

	auth.GET("/cart", controllers.GetKeranjangByIDWarga)
	auth.GET("/cart/:id", controllers.GetKeranjangByIDWarga)
	// auth.PUT("/order/:id", controllers.UpdateOrderById)
	// auth.DELETE("/order/:id", controllers.SoftDeleteOrderById)

	return e
}
