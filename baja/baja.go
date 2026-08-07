package baja

import (
	"fmt"
	"html/template"
	"net/url"
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
	Post   Post
	Posts  Posts
}

type Link struct {
	URI         string        `yaml:"url"`
	UtmURI      string        `yaml:"-"`
	Title       string        `yaml:"title"`
	Description template.HTML `yaml:"description"`
	Category    []string      `yaml:"category"`
	Action      string        `yaml:"action"`
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
	Subject string `yaml:"subject"`
	Time    string `yaml:"time"`

	Links      []Link `yaml:"links"`
	CodeToRead []Link `yaml:"read_code"`
	Tools      []Link `yaml:"tool"`
	Briefs     []Link `yaml:"brief"`
	Videos     []Link `yaml:"video"`
	SelfHosted []Link `yaml:"self_hosted"`

	Name             string          `yaml:"name"`
	PubTime          time.Time       `yaml:"pub_time"`
	Draft            bool            `yaml:"draft"`
	Description      template.HTML   `yaml:"description"`
	ExternalSnippets []template.HTML `yaml:"external_snippets"`
}

func (issue Issue) FormatPubTime() string {
	return issue.PubTime.Format("Mon, 2 Jan 2006 15:04:05 MST")
}

func (issue *Issue) Utmify(medium string) {
	for i, link := range issue.Links {
		parsedURL, err := url.Parse(link.URI)

		if err != nil {
			fmt.Println("Error parsing URL:", err)
			issue.Links[i].UtmURI = issue.Links[i].URI
			continue
		}

		// Get the existing query parameters
		queryParams := parsedURL.Query()

		// Add UTM parameters
		queryParams.Set("utm_source", "betterdev.link")
		queryParams.Set("utm_medium", medium)
		queryParams.Set("utm_campaign", fmt.Sprintf("issue-%s", issue.Name))
		parsedURL.RawQuery = queryParams.Encode()

		// Rebuild the URL with the new query parameters
		issue.Links[i].UtmURI = parsedURL.String()
	}
	for i, link := range issue.CodeToRead {
		parsedURL, err := url.Parse(link.URI)

		if err != nil {
			fmt.Println("Error parsing URL:", err)
			issue.CodeToRead[i].UtmURI = issue.CodeToRead[i].URI
			continue
		}

		// Get the existing query parameters
		queryParams := parsedURL.Query()

		// Add UTM parameters
		queryParams.Set("utm_source", "betterdev.link")
		queryParams.Set("utm_medium", medium)
		queryParams.Set("utm_campaign", fmt.Sprintf("issue-%s", issue.Name))
		parsedURL.RawQuery = queryParams.Encode()

		// Rebuild the URL with the new query parameters
		issue.CodeToRead[i].UtmURI = parsedURL.String()
	}
	for i, link := range issue.Briefs {
		parsedURL, err := url.Parse(link.URI)

		if err != nil {
			fmt.Println("Error parsing URL:", err)
			issue.Briefs[i].UtmURI = issue.Briefs[i].URI
			continue
		}

		// Get the existing query parameters
		queryParams := parsedURL.Query()

		// Add UTM parameters
		queryParams.Set("utm_source", "betterdev.link")
		queryParams.Set("utm_medium", medium)
		queryParams.Set("utm_campaign", fmt.Sprintf("issue-%s", issue.Name))
		parsedURL.RawQuery = queryParams.Encode()

		// Rebuild the URL with the new query parameters
		issue.Briefs[i].UtmURI = parsedURL.String()
	}

	for i, link := range issue.Tools {
		parsedURL, err := url.Parse(link.URI)

		if err != nil {
			fmt.Println("Error parsing URL:", err)
			issue.Tools[i].UtmURI = issue.Tools[i].URI
			continue
		}

		// Get the existing query parameters
		queryParams := parsedURL.Query()

		// Add UTM parameters
		queryParams.Set("utm_source", "betterdev.link")
		queryParams.Set("utm_medium", medium)
		queryParams.Set("utm_campaign", fmt.Sprintf("issue-%s", issue.Name))
		parsedURL.RawQuery = queryParams.Encode()

		// Rebuild the URL with the new query parameters
		issue.Tools[i].UtmURI = parsedURL.String()
	}

	for i, link := range issue.Videos {
		parsedURL, err := url.Parse(link.URI)

		if err != nil {
			fmt.Println("Error parsing URL:", err)
			issue.Videos[i].UtmURI = issue.Videos[i].URI
			continue
		}

		// Get the existing query parameters
		queryParams := parsedURL.Query()

		// Add UTM parameters
		queryParams.Set("utm_source", "betterdev.link")
		queryParams.Set("utm_medium", medium)
		queryParams.Set("utm_campaign", fmt.Sprintf("issue-%s", issue.Name))
		parsedURL.RawQuery = queryParams.Encode()

		// Rebuild the URL with the new query parameters
		issue.Videos[i].UtmURI = parsedURL.String()
	}

	for i, link := range issue.SelfHosted {
		parsedURL, err := url.Parse(link.URI)

		if err != nil {
			fmt.Println("Error parsing URL:", err)
			issue.SelfHosted[i].UtmURI = issue.SelfHosted[i].URI
			continue
		}

		// Get the existing query parameters
		queryParams := parsedURL.Query()

		// Add UTM parameters
		queryParams.Set("utm_source", "betterdev.link")
		queryParams.Set("utm_medium", medium)
		queryParams.Set("utm_campaign", fmt.Sprintf("issue-%s", issue.Name))
		parsedURL.RawQuery = queryParams.Encode()

		// Rebuild the URL with the new query parameters
		issue.SelfHosted[i].UtmURI = parsedURL.String()
	}

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
