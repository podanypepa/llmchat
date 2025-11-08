package gemini

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

// StreamSend sends a chat request to the Gemini API and returns a channel of chat responses.
func (c *Client) StreamSend(ctx context.Context, req *ChatRequest) (<-chan ChatResponse, error) {
	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	url := fmt.Sprintf("%s/v1beta/models/%s:streamGenerateContent?key=%s", c.config.BaseURL, c.config.Model, c.config.APIKey)
	httpReq, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := llmrequest.SendRequest(
		ctx,
		httpReq,
		map[string]string{
			"accept": "application/json",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	ch := make(chan ChatResponse)
	go func() {
		defer resp.Body.Close()
		defer close(ch)

		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "data:") {
				data := strings.TrimPrefix(line, "data: ")
				var chatResponse ChatResponse
				if err := json.Unmarshal([]byte(data), &chatResponse); err != nil {
					fmt.Printf("Error unmarshalling chunk: %v\n", err)
					continue
				}
				ch <- chatResponse
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Printf("Scanner error: %v\n", err)
		}
	}()

	return ch, nil
}
