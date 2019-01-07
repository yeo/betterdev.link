package server

import (
  "time"

  "github.com/labstack/echo"
  "github.com/labstack/echo/middleware"
)

// https://github.com/LYY/echo-middleware/blob/master/nocache.go
func noCache()  echo.MiddlewareFunc {
  epoch := time.Unix(0, 0).Format(time.RFC1123)

  // Taken from https://github.com/mytrile/nocache
  noCacheHeaders := map[string]string{
    "Expires":         epoch,
    "Cache-Control":   "no-cache, private, max-age=0",
    "Pragma":          "no-cache",
    "X-Accel-Expires": "0",
  }


  return func(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) (err error) {
      // Set our NoCache headers
      res := c.Response()
      for k, v := range noCacheHeaders {
        res.Header().Set(k, v)
      }

      return next(c)
    }
  }
}


func router(e *echo.Echo, s *Server) {
  e.Use(middleware.Logger())
  e.Use(noCache())

  e.POST("/githook", s.Githook)
  e.GET("/_stage/:repo", s.StageRelease)
  e.GET("/links/:url", s.VisitLink)
  e.GET("/tracks/:issue", s.Track)
  e.Static("/admin", "frontend/betterdev")
  e.GET("/api/activity/:email", s.Activity)
  e.Static("/", "public")
}
