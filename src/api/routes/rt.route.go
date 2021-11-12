package routes

import (
	"src/api/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RT(e *echo.Echo, JWTconfig middleware.JWTConfig) *echo.Echo {

	e.POST("/rt", controllers.CreateRT)
	e.GET("/rt", controllers.GetAllRT)
	e.GET("/rt/pengurus", controllers.GetAllRTWithPengurus)
	e.GET("/rt/keluarga", controllers.GetAllRTWithKeluarga)
	e.GET("/rt/:id", controllers.GetRTByID)
	e.PUT("/rt/:id", controllers.UpdateRTById)
	e.DELETE("/rt/:id", controllers.SoftDeleteRTById)

	return e
}
