package anthropic

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/podanypepa/llmchat/pkg/llmrequest"
)

// StreamSend sends a request to the Anthropic API and returns a channel of stream events.
func (c *Client) StreamSend(ctx context.Context, req *ChatRequest) (<-chan StreamEvent, error) {
	if err := validate(req); err != nil {
		return nil, fmt.Errorf("invalid anthropic chat request: %w", err)
	}
	req.Stream = true

	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
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
			"accept":            "text/event-stream",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	ch := make(chan StreamEvent)
	go func() {
		defer resp.Body.Close()
		defer close(ch)

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "data:") {
				data := strings.TrimPrefix(line, "data: ")
				var event StreamEvent
				if err := json.Unmarshal([]byte(data), &event); err != nil {
					fmt.Printf("Error unmarshalling event: %v\n", err)
					continue
				}
				ch <- event
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Scanner error: %v\n", err)
		}
	}()

	return ch, nil
}
