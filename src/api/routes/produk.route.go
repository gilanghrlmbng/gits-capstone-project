package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Produk(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {
	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(JWTconfig))
	auth.POST("/produk", controllers.CreateProduk)
	auth.GET("/produk", controllers.GetAllProduk)
	auth.GET("/produk/keluarga", controllers.GetAllProdukByKeluarga)
	auth.GET("/produk/:id", controllers.GetProdukByID)
	auth.PUT("/produk/:id", controllers.UpdateProdukById)
	auth.DELETE("/produk/:id", controllers.SoftDeleteProdukById)

	return e
}
