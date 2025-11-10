package grok

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/podanypepa/llmchat/pkg/llmrequest"
)

// Send sends a chat request to the ChatGPT API and returns the response.
func (c *Client) send(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	endpoint := "/v1/chat/completions"
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.config.BaseURL+endpoint, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := llmrequest.SendRequest(
		ctx,
		httpReq,
		map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", c.config.APIKey),
			"content-type":  "application/json",
			"accept":        "application/json",
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
