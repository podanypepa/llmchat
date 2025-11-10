package ollama

import (
	"cmp"
	"fmt"
)

// Client is the Ollama API client.
type Client struct {
	config *Config
}

// Config holds the configuration for the Client.
type Config struct {
	BaseURL string
}

// NewClient creates a new Client.
func NewClient() (*Client, error) {
	return &Client{
		config: &Config{
			BaseURL: DefaultBaseURL,
		},
	}, nil
}

// NewClientWithConfig creates a new Client with the provided configuration.
func NewClientWithConfig(config *Config) (*Client, error) {
	if config == nil {
		return nil, fmt.Errorf("config is required")
	}

	config.BaseURL = cmp.Or(config.BaseURL, DefaultBaseURL)

	return &Client{
		config: config,
	}, nil
}