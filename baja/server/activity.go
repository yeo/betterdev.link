package server

import (
	"net/http"

	"github.com/yeo/betterdev.link/baja/dts"

	"github.com/labstack/echo"
)

// Track returns a 1x1 transparent gif pixel for tracking purpose
func (s *Server) Activity(c echo.Context) error {
	user := c.Param("email")

	activity := dts.ActivityService{s.db}
	userHistory := activity.LoadUser(user)

	c.Response().Header().Set(echo.HeaderContentType, "application/json")

	return c.JSON(http.StatusOK, userHistory)
}
