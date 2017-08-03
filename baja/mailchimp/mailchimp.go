package mailchimp

import (
	"net/http"
	"os"
	"time"
)

type Client struct {
	endpoint string
	client   *http.Client
}

func GetClient() *Client {
	c := &Client{
		endpoint: os.Getenv("MAILCHIMP_API_URI"),
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}

	return c
}

type Campaign struct {
	ID    string
	WebID string
}

func NewCampain(type_ string) {

}
