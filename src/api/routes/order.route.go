package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Order(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {
	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(JWTconfig))

	auth.POST("/order/create/:pembayaran", controllers.CreateOrder)
	auth.PUT("/order/proses/:id", controllers.OrderProses)
	auth.PUT("/order/cancel/:id", controllers.OrderCancel)
	auth.PUT("/order/selesai/:id", controllers.OrderSelesai)

	auth.GET("/order/warga", controllers.GetAllOrderPembeli)
	auth.GET("/order/warga/:id", controllers.GetOrderByIdPembeli)
	auth.GET("/order/toko", controllers.GetAllOrderPenjual)
	auth.GET("/order/toko/:id", controllers.GetOrderByIdPenjual)
	auth.GET("/order/:id", controllers.GetOrderByID)
	// auth.PUT("/order/:id", controllers.UpdateOrderById)
	// auth.DELETE("/order/:id", controllers.SoftDeleteOrderById)

	return e
}
