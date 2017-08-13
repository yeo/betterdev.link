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
	URI         string   `yaml:"url"`
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
	Category    []string `yaml:"category"`
	Action      string   `yaml:"action"`
}

func (l *Link) IsSponsor() bool {
	for _, c := range l.Category {
		if c == "sponsor" {
			return true
		}
	}

	return false
}

type Issues []Issue

type Issue struct {
	Time        string    `yaml:"time"`
	Links       []Link    `yaml:"links"`
	CodeToRead  []Link    `yaml:"read_code"`
	Tools       []Link    `yaml:"tool"`
	Name        string    `yaml:"name"`
	PubTime     time.Time `yaml:"pub_time"`
	Draft       bool      `yaml:"draft"`
	Description string    `yaml:"description"`
}

func (issue Issue) FormatPubTime() string {
	return issue.PubTime.Format("Mon, 2 Jan 2006 15:04:05 MST")
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
