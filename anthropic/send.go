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
func (c *Client) Send(ctx context.Context, req *Request) (*Response, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	if jsonData == nil {
		return nil, fmt.Errorf("request is empty")
	}

	endpoint := "/v1/messages"
	httpReq, err := http.NewRequest("POST", c.config.BaseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	httpReq = httpReq.WithContext(ctx)

	resultChan := make(chan struct {
		resp *http.Response
		err  error
	}, 1)

	go func() {
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

		var chatResponse Response
		if err := json.NewDecoder(result.resp.Body).Decode(&chatResponse); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &chatResponse, nil
	}
}
