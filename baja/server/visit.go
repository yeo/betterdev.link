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

	"github.com/yeo/betterdev.link/baja/dts"
)

// Allow us to view any branch of code from any accessible git repository
func (s *Server) VisitLink(c echo.Context) error {
	url64 := c.Param("url")
	url, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(url64)

	if err != nil {
		return c.HTML(http.StatusNotAcceptable, "Invalid link")
	}

	link := string(url)

	tracker := dts.TrackerService{s.db}
	// TODO: Need input validation for issue/email and sanitize them
	tracker.OpenURL(link, c.QueryParam("issue"), c.QueryParam("email"), string(c.RealIP()))

	return c.Redirect(http.StatusTemporaryRedirect, link)
}
