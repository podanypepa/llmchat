package chatgpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/podanypepa/llmchat/pkg/llmrequest"
)

// Send sends a chat request to the ChatGPT API and returns the response.
func (c *Client) Send(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	if err := validate(req); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	endpoint := "/chat/completions"
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

	var chatResponse ChatResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &chatResponse, nil
}
