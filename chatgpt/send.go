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
	if reqBytes == nil {
		return nil, fmt.Errorf("request is empty")
	}

	endpoint := "/chat/completions"
	httpReq, err := http.NewRequest("POST", c.config.BaseURL+endpoint, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resultChan := make(chan struct {
		resp *http.Response
		err  error
	}, 1)

	go func() {
		headders := map[string]string{
			"Authorization": fmt.Sprintf("Bearer %s", c.config.APIKey),
			"Content-Type":  "application/json",
			"Accept":        "application/json",
		}
		if c.config.OrganizationID != "" {
			headders["OpenAI-Organization"] = c.config.OrganizationID
		}
		resp, err := llmrequest.SendRequest(
			ctx,
			httpReq,
			headders,
		)

		resultChan <- struct {
			resp *http.Response
			err  error
		}{resp, err}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case result := <-resultChan:
		defer result.resp.Body.Close()
		if result.err != nil {
			return nil, fmt.Errorf("request failed: %w", result.err)
		}

		var chatResponse ChatResponse
		if err := json.NewDecoder(result.resp.Body).Decode(&chatResponse); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &chatResponse, nil
	}
}
