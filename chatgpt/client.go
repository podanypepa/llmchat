package chatgpt

import (
	"cmp"
	"fmt"
)

// Client is the ChatGPT API client.
type Client struct {
	config *Config
}

// Config holds the configuration for the Client.
type Config struct {
	BaseURL        string
	APIKey         string
	OrganizationID string
}

// NewClient creates a new Client with the provided API key.
func NewClient(apikey string) (*Client, error) {
	if apikey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	return &Client{
		config: &Config{
			BaseURL: DefaultapiURL,
			APIKey:  apikey,
		},
	}, nil
}

// NewClientWithConfig creates a new Client with the provided configuration.
func NewClientWithConfig(config *Config) (*Client, error) {
	if config.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	config.BaseURL = cmp.Or(config.BaseURL, DefaultapiURL)

	return &Client{
		config: config,
	}, nil
}
