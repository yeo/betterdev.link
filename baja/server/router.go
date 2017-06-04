package server

import (
	"github.com/labstack/echo"
)

func router(e *echo.Echo) {
	e.Static("/", "public")
}
