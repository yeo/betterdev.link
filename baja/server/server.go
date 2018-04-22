package server

import (
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

	router(e, s)
	e.Logger.Fatal(e.Start(addr))
}
