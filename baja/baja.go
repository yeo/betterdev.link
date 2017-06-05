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
	Time  time.Time
	Issue Issue
}

type Link struct {
	URI   string
	Title string
}

type Issue struct {
	Name  string
	Time  time.Time
	Links []Link
}
