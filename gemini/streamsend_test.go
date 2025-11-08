package gemini

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStreamSend(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// A sample stream of events
		data := `
data: {"candidates":[{"content":{"parts":[{"text":"Hello"}],"role":"model"},"finishReason":"STOP","index":0,"safetyRatings":[{"category":"HARM_CATEGORY_SEXUALLY_EXPLICIT","probability":"NEGLIGIBLE"},{"category":"HARM_CATEGORY_HATE_SPEECH","probability":"NEGLIGIBLE"},{"category":"HARM_CATEGORY_HARASSMENT","probability":"NEGLIGIBLE"},{"category":"HARM_CATEGORY_DANGEROUS_CONTENT","probability":"NEGLIGIBLE"}]}]}

data: {"candidates":[{"content":{"parts":[{"text":" world"}],"role":"model"},"finishReason":"STOP","index":0,"safetyRatings":[{"category":"HARM_CATEGORY_SEXUALLY_EXPLICIT","probability":"NEGLIGIBLE"},{"category":"HARM_CATEGORY_HATE_SPEECH","probability":"NEGLIGIBLE"},{"category":"HARM_CATEGORY_HARASSMENT","probability":"NEGLIGIBLE"},{"category":"HARM_CATEGORY_DANGEROUS_CONTENT","probability":"NEGLIGIBLE"}]}]}
`
		fmt.Fprint(w, data)
	}))
	defer server.Close()

	config := &Config{
		BaseURL: server.URL,
		APIKey:  "test-key",
		Model:   "gemini-pro",
	}
	client, err := NewClientWithConfig(config)
	if err != nil {
		t.Fatalf("NewClientWithConfig() error = %v", err)
	}

	req := &ChatRequest{
		Contents: []Content{
			{
				Role: "user",
				Parts: []Part{
					{Text: "Hello"},
				},
			},
		},
	}

	ch, err := client.StreamSend(context.Background(), req)
	if err != nil {
		t.Fatalf("StreamSend() error = %v", err)
	}

	var responses []ChatResponse
	for resp := range ch {
		responses = append(responses, resp)
	}

	if len(responses) != 2 {
		t.Errorf("Expected 2 responses, got %d", len(responses))
	}

	expectedTexts := []string{"Hello", " world"}
	for i, resp := range responses {
		text, err := resp.ExtractText()
		if err != nil {
			t.Errorf("Error extracting text from response %d: %v", i, err)
		}
		if text != expectedTexts[i] {
			t.Errorf("Expected text '%s', got '%s'", expectedTexts[i], text)
		}
	}
}
