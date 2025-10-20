package gemini

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
		return nil, fmt.Errorf("model is required")
	}

	return &Client{
		config: &Config{
			APIKey:       apikey,
			DefaultModel: DefaultModel,
			Model:        model,
		},
	}, nil
}

// NewClientWithConfig creates a new Client with the provided configuration.
func NewClientWithConfig(config *Config) (*Client, error) {
	if config == nil {
		return nil, fmt.Errorf("config is required")
	}
	if config.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}
	if config.Model == "" {
		return nil, fmt.Errorf("model is required")
	}

	config.DefaultModel = cmp.Or(config.DefaultModel, DefaultModel)

	return &Client{
		config: config,
	}, nil
}
