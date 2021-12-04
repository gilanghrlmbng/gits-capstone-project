package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Informasi(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {
	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(JWTconfig))
	auth.POST("/informasi", controllers.CreateInformasi)
	auth.GET("/informasi", controllers.GetAllInformasi)
	auth.GET("/informasi/:id", controllers.GetInformasiByID)
	auth.GET("/infoterkini", controllers.GetInfoTerkini)
	auth.GET("/kegiatan", controllers.GetKegiatanWarga)
	auth.PUT("/informasi/:id", controllers.UpdateInformasiById)
	auth.DELETE("/informasi/:id", controllers.SoftDeleteInformasiById)

	return e
}
