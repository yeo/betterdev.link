package server

import (
	"github.com/labstack/echo"
	"net/http"
)

func Run(addr string) {
	e := echo.New()
	router(e)
	e.Logger.Fatal(e.Start(addr))
}
