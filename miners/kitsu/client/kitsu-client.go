package kitsuclient

import "net/http"

// Kitsu is a really dumb API client, only implementing relevant parts of
// the kitsu API for our usage
type Kitsu struct {
	c       *http.Client
	baseURL string
}

// NewClient returns a Kitsu client ready for use.
func NewClient() (*Kitsu, error) {
	return &Kitsu{
		c:       &http.Client{},
		baseURL: "https://kitsu.io/api/edge",
	}, nil
}
