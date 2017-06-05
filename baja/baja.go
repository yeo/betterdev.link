package baja

import (
	"strconv"
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
	Issues Issues
}

type Link struct {
	URI         string `yaml:"url"`
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
}

type Issues []Issue

type Issue struct {
	Time    string    `yaml:"time"`
	Links   []Link    `yaml:"links"`
	Name    string    `yaml:"name"`
	PubTime time.Time `yaml:"pub_time"`
}

func (issue Issues) Len() int {
	return len(issue)
}

func (issue Issues) Less(i, j int) bool {
	i1, err := strconv.ParseInt(issue[i].Name, 10, 32)
	if err != nil {
		i1 = 0
	}

	j1, err := strconv.ParseInt(issue[j].Name, 10, 32)
	if err != nil {
		j1 = 0
	}

	return i1 < j1
}

func (issue Issues) Swap(ii, j int) {
	issue[ii], issue[j] = issue[j], issue[ii]
}
