// Package anthropic ...
package anthropic

import "net/http"

// Client of the Anthropic API.
type Client struct {
	apiKey     string
	httpClient *http.Client
}

// NewClient creates a new Anthropic API client with the given API key.
func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}
