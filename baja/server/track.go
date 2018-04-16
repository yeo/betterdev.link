package server

import (
	"encoding/base64"
	"net/http"

	"github.com/yeo/betterdev.link/baja/dts"

	"github.com/labstack/echo"
)

const base64GifPixel = "R0lGODlhAQABAIAAAP///wAAACwAAAAAAQABAAACAkQBADs="

// Track returns a 1x1 transparent gif pixel for tracking purpose
func (s *Server) Track(c echo.Context) error {
	issue := c.Param("issue")

	tracker := dts.TrackerService{s.db}
	tracker.LoadIssue(issue, c.QueryParam("email"), string(c.RealIP()))

	c.Response().Header().Set(echo.HeaderContentType, "image/gif")
	img, _ := base64.StdEncoding.DecodeString(base64GifPixel)

	return c.String(http.StatusOK, string(img))
}
