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
	log "github.com/sirupsen/logrus"
)

const base64GifPixel = "R0lGODlhAQABAIAAAP///wAAACwAAAAAAQABAAACAkQBADs="

// Allow us to view any branch of code from any accessible git repository
func (s *Server) Track(c echo.Context) error {
	issue := c.Param("issue")

	s.log.WithFields(log.Fields{
		"issue": issue,
		"email": c.QueryParam("email"),
		"ip":    string(c.RealIP()),
	}).Info("Open Issue")

	c.Response().Header().Set(echo.HeaderContentType, "image/gif")
	img, _ := base64.StdEncoding.DecodeString(base64GifPixel)

	return c.String(http.StatusOK, string(img))
}
