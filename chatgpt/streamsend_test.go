package chatgpt

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStreamSend(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.WriteHeader(http.StatusOK)
		// A sample stream of events
		data := `data: {"id":"chatcmpl-123","object":"chat.completion.chunk","created":1694268190,"model":"gpt-4o-mini","choices":[{"index":0,"delta":{"role":"assistant","content":""},"finish_reason":null}]}

data: {"id":"chatcmpl-123","object":"chat.completion.chunk","created":1694268190,"model":"gpt-4o-mini","choices":[{"index":0,"delta":{"content":"Hello"},"finish_reason":null}]}

data: {"id":"chatcmpl-123","object":"chat.completion.chunk","created":1694268190,"model":"gpt-4o-mini","choices":[{"index":0,"delta":{"content":" world"},"finish_reason":null}]}

data: {"id":"chatcmpl-123","object":"chat.completion.chunk","created":1694268190,"model":"gpt-4o-mini","choices":[{"index":0,"delta":{},"finish_reason":"stop"}]}

data: [DONE]
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
		Model: Gpt4OMini,
		Messages: []ChatMessage{
			{
				Role:    RoleUser,
				Content: "Hello",
			},
		},
	}

	ch, err := client.StreamSend(context.Background(), req)
	if err != nil {
		t.Fatalf("StreamSend() error = %v", err)
	}

	var chunks []StreamCompletionChunk
	for chunk := range ch {
		chunks = append(chunks, chunk)
	}

	if len(chunks) != 4 {
		t.Errorf("Expected 4 chunks, got %d", len(chunks))
	}

	expectedContents := []string{"", "Hello", " world", ""}
	for i, chunk := range chunks {
		if len(chunk.Choices) > 0 {
			if chunk.Choices[0].Delta.Content != expectedContents[i] {
				t.Errorf("Expected content '%s', got '%s'", expectedContents[i], chunk.Choices[0].Delta.Content)
			}
		}
	}
}
