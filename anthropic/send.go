package anthropic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/podanypepa/llmchat/pkg/llmrequest"
)

// Send sends a request to the Anthropic API and returns the response.
func (c *Client) Send(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	if err := validate(req); err != nil {
		return nil, fmt.Errorf("invalid anthropic chat request: %w", err)
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	if jsonData == nil {
		return nil, fmt.Errorf("request is empty")
	}

	endpoint := "/v1/messages"
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.config.BaseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	resp, err := llmrequest.SendRequest(
		ctx,
		httpReq,
		map[string]string{
			"x-api-key":         c.apiKey,
			"content-type":      "application/json",
			"anthropic-version": "2023-06-01",
			"accept":            "application/json",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var chatResponse ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &chatResponse, nil
}
