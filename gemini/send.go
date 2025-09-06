package gemini

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/podanypepa/llmchat/pkg/llmrequest"
)

// Send sends a chat request to the ChatGPT API and returns the response.
func (c *Client) Send(ctx context.Context, req *GenerateContentRequest) (*GenerateContentResponse, error) {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}
	if reqBytes == nil {
		return nil, fmt.Errorf("request is empty")
	}

	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", c.config.Model, c.config.APIKey)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	httpReq = httpReq.WithContext(ctx)

	resultChan := make(chan struct {
		resp *http.Response
		err  error
	}, 1)

	go func() {
		resp, err := llmrequest.SendRequest(ctx, httpReq, map[string]string{"accept": "application/json"})
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

		var chatResponse GenerateContentResponse
		if err := json.NewDecoder(result.resp.Body).Decode(&chatResponse); err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}
		return &chatResponse, nil
	}
}
