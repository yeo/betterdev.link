package server

import (
	"github.com/labstack/echo"
)

func Run(addr string) {
	e := echo.New()
	router(e)
	e.Logger.Fatal(e.Start(addr))
}
