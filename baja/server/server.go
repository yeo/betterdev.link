package server

import (
	"github.com/labstack/echo"
)

type Server struct{}

func Run(addr string) {
	e := echo.New()
	s := &Server{}
	router(e, s)
	e.Logger.Fatal(e.Start(addr))
}
