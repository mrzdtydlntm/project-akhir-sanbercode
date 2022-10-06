package routes

import (
	"sanbertutor/controller"
	sanberMiddleware "sanbertutor/middleware"

	"github.com/labstack/echo"
)

func Routes() *echo.Echo {
	e := echo.New()

	sanberMiddleware.LoggerMiddleware(e)

	e.POST("/register", controller.RegisterController)
	e.POST("/login", controller.LoginController)

	return e
}
