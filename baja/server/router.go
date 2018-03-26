package server

import (
	"github.com/labstack/echo"
)

func router(e *echo.Echo, s *Server) {
	e.POST("/githook", s.Githook)
	e.GET("/_stage/:repo", s.StageRelease)
	e.GET("/links/:url", s.VisitLink)
	e.Static("/", "public")
}
