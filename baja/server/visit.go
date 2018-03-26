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

// Allow us to view any branch of code from any accessible git repository
func (s *Server) VisitLink(c echo.Context) error {
	url64 := c.Param("url")
	url, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(url64)

	if err != nil {
		return c.HTML(http.StatusNotAcceptable, "Invalid link")
	}

	link := string(url)

	s.log.WithFields(log.Fields{
		"url":   link,
		"email": c.QueryParam("email"),
		"ip":    string(c.RealIP()),
	}).Info("Open URL")

	return c.Redirect(http.StatusTemporaryRedirect, link)
}
