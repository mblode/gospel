package routes

import (
	"github.com/mblode/gospel/app/controllers"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", controllers.Home)

	e.GET("/users", controllers.GetUsers)
	return e
}
