package routes

import (
	"acp8/constants"
	"acp8/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	eJWT := e.Group("")
	eJWT.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	eJWT.GET("/users", controllers.GetUserControllers)
	eJWT.POST("/registerProduct", controllers.RegisterProductController)

	// e.POST("/login", loginControllers)
	e.POST("/register", controllers.RegisterControllers)
	e.GET("/products/:id", controllers.GetProductWIthParamsController)
	return e
}
