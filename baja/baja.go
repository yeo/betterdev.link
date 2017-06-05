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

// Page holds the whole state of current page, including config, site etc
type Page struct {
	Site   Site
	Time   time.Time
	Issue  Issue
	Issues []Issue
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
