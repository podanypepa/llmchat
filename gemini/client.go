package gemini

import (
	"cmp"
	"fmt"
	"net/http"
)

// Client is the ChatGPT API client.
type Client struct {
	client *http.Client
	config *Config
}

// Config holds the configuration for the Client.
type Config struct {
	APIKey       string
	Model        string
	DefaultModel string
}

// NewClient creates a new Client with the provided API key.
func NewClient(apikey string, model string) (*Client, error) {
	if apikey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	if model == "" {
		return nil, fmt.Errorf("Model is required")
	}

	return &Client{
		client: &http.Client{},
		config: &Config{
			APIKey:       apikey,
			DefaultModel: DefaultModel,
			Model:        model,
		},
	}, nil
}

// NewClientWithConfig creates a new Client with the provided configuration.
func NewClientWithConfig(config *Config) (*Client, error) {
	if config.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}
	if config.Model == "" {
		return nil, fmt.Errorf("Model is required")
	}

	config.DefaultModel = cmp.Or(config.DefaultModel, DefaultModel)

	return &Client{
		client: &http.Client{},
		config: config,
	}, nil
}
