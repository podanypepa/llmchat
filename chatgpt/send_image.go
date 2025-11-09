package chatgpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/podanypepa/llmchat/pkg/llmrequest"
)

// SendImageRequest sends an image generation request to the DALL-E API and returns the response.
func (c *Client) SendImageRequest(ctx context.Context, req ImageRequest) (*ImageResponse, error) {
	// Validation for image request can be added here if needed.

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	endpoint := "/images/generations"
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.config.BaseURL+endpoint, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", c.config.APIKey),
		"Content-Type":  "application/json",
		"Accept":        "application/json",
	}
	if c.config.OrganizationID != "" {
		headers["OpenAI-Organization"] = c.config.OrganizationID
	}

	resp, err := llmrequest.SendRequest(ctx, httpReq, headers)
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
