package anthropic

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStreamSend(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		// A sample stream of events from the Anthropic API documentation
		data := `event: message_start
data: {"type": "message_start", "message": {"id": "msg_123", "type": "message", "role": "assistant", "model": "claude-3-opus-20240229", "content": [], "stop_reason": null, "stop_sequence": null, "usage": {"input_tokens": 10, "output_tokens": 1}}}

event: content_block_start
data: {"type": "content_block_start", "index": 0, "content_block": {"type": "text", "text": ""}}

event: content_block_delta
data: {"type": "content_block_delta", "index": 0, "delta": {"type": "text_delta", "text": "Hello"}}

event: content_block_delta
data: {"type": "content_block_delta", "index": 0, "delta": {"type": "text_delta", "text": " world"}}

event: content_block_stop
data: {"type": "content_block_stop", "index": 0}

event: message_delta
data: {"type": "message_delta", "delta": {"stop_reason": "end_turn", "stop_sequence":null, "usage":{"output_tokens": 15}}}

event: message_stop
data: {"type": "message_stop"}
`
		w.Write([]byte(data))
	}))
	defer server.Close()

	config := &Config{
		BaseURL: server.URL,
		APIKey:  "test-key",
	}
	client, err := NewClientWithConfig(config)
	if err != nil {
		t.Fatalf("NewClientWithConfig() error = %v", err)
	}

	req := &ChatRequest{
		Model: "claude-3-opus-20240229",
		Messages: []Message{
			{
				Role:    RoleUser,
				Content: "Hello",
			},
		},
		MaxTokens: 1024,
	}

	ch, err := client.StreamSend(context.Background(), req)
	if err != nil {
		t.Fatalf("StreamSend() error = %v", err)
	}

	var events []StreamEvent
	for event := range ch {
		events = append(events, event)
	}

	if len(events) != 7 {
		t.Errorf("Expected 7 events, got %d", len(events))
	}

	expectedTexts := []string{"Hello", " world"}
	textIndex := 0
	for _, event := range events {
		if event.Type == "content_block_delta" {
			if event.Delta.Text != expectedTexts[textIndex] {
				t.Errorf("Expected text '%s', got '%s'", expectedTexts[textIndex], event.Delta.Text)
			}
			textIndex++
		}
	}
}
