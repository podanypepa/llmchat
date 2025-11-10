package ollama

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/podanypepa/llmchat/pkg/llmrequest"
)

// StreamSend sends a chat request to the Ollama API and returns a channel of stream completion chunks.
func (c *Client) StreamSend(ctx context.Context, req *ChatRequest) (<-chan StreamCompletionChunk, error) {
	req.Stream = true
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	endpoint := "/api/chat"
	httpReq, err := http.NewRequestWithContext(ctx, "POST", c.config.BaseURL+endpoint, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/x-ndjson",
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
			line := scanner.Bytes()
			if len(line) == 0 {
				continue
			}

			var chunk StreamCompletionChunk
			if err := json.Unmarshal(line, &chunk); err != nil {
				fmt.Printf("Error unmarshalling chunk: %v\n", err)
				continue
			}
			ch <- chunk
			if chunk.Done {
				return
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Scanner error: %v\n", err)
		}
	}()

	return ch, nil
}