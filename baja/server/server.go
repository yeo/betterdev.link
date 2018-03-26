package server

import (
	"os"

	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	log *log.Logger
}

func Run(addr string) {
	e := echo.New()
	s := &Server{
		log: log.New(),
	}

	file, err := os.OpenFile("/var/log/bd/click.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Cannot open log file")
	}

	s.log.Out = file
	router(e, s)
	e.Logger.Fatal(e.Start(addr))
}
