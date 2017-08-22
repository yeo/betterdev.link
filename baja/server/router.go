package server

import (
	"github.com/labstack/echo"
)

func router(e *echo.Echo, s *Server) {
	e.POST("/githook", s.Githook)
	e.Static("/", "public")
}
