package chatgpt

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Responses sends a request to the Responses API and returns the response.
// API specification: api.openai.com/v1/responses
func (c *Client) Responses(
	ctx context.Context,
	req *ResponsesRequest,
) (
	*ResponsesResponse,
	error,
) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Responses: marshal request error: %w", err)
	}

	httpReq, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		"https://api.openai.com/v1/responses",
		bytes.NewReader(b),
	)
	if err != nil {
		return nil, fmt.Errorf("Responses: create http request error: %w", err)
	}

	httpReq.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		var buf bytes.Buffer
		_, _ = buf.ReadFrom(resp.Body)
		return nil, fmt.Errorf("Responses: non-2xx status code: %s", resp.Status)
	}

	resByte, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Responses: read response body error: %w", err)
	}

	var respData ResponsesResponse
	if err := json.Unmarshal(resByte, &respData); err != nil {
		return nil, fmt.Errorf("Responses: decode response error: %w", err)
	}

	if len(respData.Output) == 0 {
		fmt.Println("No output from model.")
		return nil, fmt.Errorf("Responses: empty output from model")
	}

	return &respData, nil
}
