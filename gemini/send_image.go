package gemini

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/podanypepa/llmchat/pkg/llmrequest"
)

// SendImageRequest sends an image generation request to the Gemini API and returns the response.
func (c *Client) SendImageRequest(ctx context.Context, req ImageRequest) (*ImageResponse, error) {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	// Note: The model for image generation is often specified in the URL.
	// Here we assume a conventional model name like "imagen-4" but this might need configuration.
	url := fmt.Sprintf("%s/v1beta/models/imagen-4:generateImages?key=%s", c.config.BaseURL, c.config.APIKey)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := llmrequest.SendRequest(
		ctx,
		httpReq,
		map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var imageResponse ImageResponse
	if err := json.NewDecoder(resp.Body).Decode(&imageResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &imageResponse, nil
}
