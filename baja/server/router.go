package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func router(e *echo.Echo, s *Server) {
	e.Use(middleware.Logger())
	e.POST("/githook", s.Githook)
	e.GET("/_stage/:repo", s.StageRelease)
	e.GET("/links/:url", s.VisitLink)
	e.Static("/", "public")
}
