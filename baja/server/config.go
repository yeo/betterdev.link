package server

import (
	"os"
)

type Config struct {
	MongoURI string
}

func LoadConfigFromEnv() *Config {
	c := Config{}

	if v := os.Getenv("MONGO_URI"); v != "" {
		c.MongoURI = v
	} else {
		c.MongoURI = "mongodb://127.0.0.1:27017"
	}

	return &c
}
