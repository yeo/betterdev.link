package server

import (
	"github.com/labstack/echo"
)

func router(e *echo.Echo) {
	e.POST("/issues", controller.index)
	e.GET("/users/:id", controller.show)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
}
