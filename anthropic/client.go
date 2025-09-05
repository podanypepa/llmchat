// Package anthropic ...
package anthropic

import (
	"cmp"
	"fmt"
	"net/http"
	"time"
)

const (
	// DefaultTimeout is the default timeout for HTTP requests.
	DefaultTimeout = 60 * time.Second
	// DefaultBaseURL is the default base URL for the Anthropic API.
	DefaultBaseURL = "https://api.anthropic.com"
)

// Config for the Anthropic API client.
type Config struct {
	APIKey      string
	BaseURL     string
	HTTPTimeout time.Duration
}

// Client of the Anthropic API.
type Client struct {
	apiKey     string
	config     Config
	httpClient *http.Client
}

// NewClient creates a new Anthropic API client with the given API key.
func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	return &Client{
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}, nil
}

func NewClientWithConfig(config *Config) (*Client, error) {
	if config.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	return &Client{
		apiKey: config.APIKey,
		httpClient: &http.Client{
			Timeout: cmp.Or(config.HTTPTimeout, DefaultTimeout),
		},
	}, nil
}
