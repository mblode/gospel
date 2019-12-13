package routes

import (
	"github.com/mblode/gospel/app/controllers"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", controllers.Home)

	e.GET("/users", controllers.GetAllUsers)
	e.GET("/users/:id", controllers.GetUser)
	e.POST("/users", controllers.NewUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
	e.PUT("/users/:id", controllers.UpdateUser)

	return e
}
