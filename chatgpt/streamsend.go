package chatgpt

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/podanypepa/llmchat/pkg/llmrequest"
)

// StreamSend sends a chat request to the ChatGPT API and returns a channel of stream completion chunks.
func (c *Client) streamSend(ctx context.Context, req *ChatRequest) (<-chan StreamCompletionChunk, error) {
	if err := validate(req); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}
	req.Stream = true

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
		"Accept":        "text/event-stream",
	}
	if c.config.OrganizationID != "" {
		headers["OpenAI-Organization"] = c.config.OrganizationID
	}

	resp, err := llmrequest.SendRequest(ctx, httpReq, headers)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	ch := make(chan StreamCompletionChunk)
	go func() {
		defer resp.Body.Close()
		defer close(ch)

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) > 6 && line[:5] == "data:" {
				if line[7:] == "[DONE]" {
					return
				}
				var chunk StreamCompletionChunk
				if err := json.Unmarshal([]byte(line[6:]), &chunk); err != nil {
					fmt.Printf("Error unmarshalling chunk: %v\n", err)
					continue
				}
				ch <- chunk
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Scanner error: %v\n", err)
		}
	}()

	return ch, nil
}
