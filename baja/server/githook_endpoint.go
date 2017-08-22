package server

import (
	"github.com/labstack/echo"
	"net/http"
)

func (s *Server) Githook(c echo.Context) error {
	//Parse and validate webhook
	// Invoke build process
	return c.String(http.StatusOK, "Hello, World!")
}
