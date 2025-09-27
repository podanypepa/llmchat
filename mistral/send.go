package mistral

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/podanypepa/llmchat/pkg/llmrequest"
)

// Send sends a chat completion request to the Mistral API and returns the response.
func (c *Client) Send(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	if reqBytes == nil {
		return nil, fmt.Errorf("request is empty")
	}

	endpoint := "/v1/chat/completions"
	httpReq, err := http.NewRequest("POST", c.config.BaseURL+endpoint, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resultChan := make(chan struct {
		resp *http.Response
		err  error
	}, 1)

	go func() {
		resp, err := llmrequest.SendRequest(
			ctx,
			httpReq,
			map[string]string{
				"Authorization": fmt.Sprintf("Bearer %s", c.config.APIKey),
				"accept":        "application/json",
				"content-type":  "application/json",
			},
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
		if result.err != nil {
			return nil, fmt.Errorf("request failed: %w", result.err)
		}
		defer result.resp.Body.Close()

		var chatResponse ChatResponse
		if err := json.NewDecoder(result.resp.Body).Decode(&chatResponse); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &chatResponse, nil
	}
}
