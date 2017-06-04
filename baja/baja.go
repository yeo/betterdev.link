package baja

import (
	"time"
)

type Config struct {
	Theme string
}

type Site struct {
	Name string
}

type Page struct {
	Time time.Time
}
