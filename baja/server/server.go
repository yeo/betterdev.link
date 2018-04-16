package server

import (
	"os"

	"github.com/labstack/echo"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
	"github.com/yeo/betterdev.link/baja/dao"
)

type Server struct {
	log    *log.Logger
	config *Config
	db     *mongo.Database
}

func Run(addr string) {
	config := LoadConfigFromEnv()
	e := echo.New()
	s := &Server{
		log:    log.New(),
		config: config,
		db:     dao.Connect(config.MongoURI),
	}

	file, err := os.OpenFile("/var/log/bd/click.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Cannot open log file")
	}

	s.log.Out = file
	router(e, s)
	e.Logger.Fatal(e.Start(addr))
}
