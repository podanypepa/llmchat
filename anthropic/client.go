// Package anthropic ...
package anthropic

import (
	"fmt"
)

const (
	// DefaultBaseURL is the default base URL for the Anthropic API.
	DefaultBaseURL = "https://api.anthropic.com"
)

// Config for the Anthropic API client.
type Config struct {
	APIKey  string
	BaseURL string
}

// Client of the Anthropic API.
type Client struct {
	apiKey string
	config *Config
}

// NewClient creates a new Anthropic API client with the given API key.
func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	return &Client{
		apiKey: apiKey,
		config: &Config{
			APIKey:  apiKey,
			BaseURL: DefaultBaseURL,
		},
	}, nil
}

// NewClientWithConfig creates a new Anthropic API client with the given configuration.
func NewClientWithConfig(config *Config) (*Client, error) {
	if config == nil {
		return nil, fmt.Errorf("config is required")
	}
	if config.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}
	if config.BaseURL == "" {
		config.BaseURL = DefaultBaseURL
	}

	return &Client{
		apiKey: config.APIKey,
		config: config,
	}, nil
}
