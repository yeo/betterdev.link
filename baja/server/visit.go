package server

import (
	"encoding/base64"
	"net/http"
	// "errors"
	// "fmt"
	// "os"
	// "strings"
	// "sync"

	"github.com/labstack/echo"
)

// Allow us to view any branch of code from any accessible git repository
func (s *Server) VisitLink(c echo.Context) error {
	url64 := c.Param("url")
	url, err := base64.StdEncoding.DecodeString(url64)

	if err == nil {
		return c.HTML(http.StatusNotAcceptable, "Invalid link")
	}

	return c.Redirect(http.StatusTemporaryRedirect, string(url))
}
