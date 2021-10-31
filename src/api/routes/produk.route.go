package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
)

func Produk(e *echo.Echo) *echo.Echo {
	e.POST("/produk", controllers.CreateProduk)
	e.GET("/produk", controllers.GetAllProduk)
	e.GET("/produk/:id", controllers.GetProdukByID)
	e.PUT("/produk/:id", controllers.UpdateProdukById)
	e.DELETE("/produk/:id", controllers.SoftDeleteProdukById)

	return e
}
