package chatgpt

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Responses(t *testing.T) {
	// Mock server
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify request method and path
		if r.Method != "POST" {
			t.Errorf("expected POST request, got %s", r.Method)
		}
		if r.URL.Path != "/responses" { // Note: BaseURL is usually "https://api.openai.com/v1", so path in mock should probably be checked relative to how client is configured.
			// However, the client appends endpoint to BaseURL.
			// If we set BaseURL to the mock server URL, the path will be just /responses.
			t.Errorf("expected /responses path, got %s", r.URL.Path)
		}

		// Verify headers
		if r.Header.Get("Authorization") != "Bearer test-api-key" {
			t.Errorf("expected Authorization header, got %s", r.Header.Get("Authorization"))
		}

		// Decode request body
		var req ResponsesRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			t.Errorf("failed to decode request body: %v", err)
		}
		if req.Model != "gpt-4o" {
			t.Errorf("expected model gpt-4o, got %s", req.Model)
		}

		// Send mock response
		resp := ResponsesResponse{
			ID:      "resp_123",
			Object:  "response",
			Created: 1234567890,
			Output: ResponsesOutput{
				Content: "Hello from mock",
			},
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			t.Errorf("failed to encode response: %v", err)
		}
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	// Create client with mock server URL
	config := &Config{
		BaseURL: server.URL, // Client uses BaseURL + endpoint. If BaseURL is "http://ip:port", endpoint "/responses" -> "http://ip:port/responses"
		APIKey:  "test-api-key",
	}
	client, err := NewClientWithConfig(config)
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
	}

	// Call Responses
	req := &ResponsesRequest{
		Model: "gpt-4o",
		Input: "Hello",
	}
	resp, err := client.Responses(context.Background(), req)
	if err != nil {
		t.Fatalf("Responses() error = %v", err)
	}

	// Verify response
	if resp.ID != "resp_123" {
		t.Errorf("expected ID resp_123, got %s", resp.ID)
	}
	if resp.Output.Content != "Hello from mock" {
		t.Errorf("expected content 'Hello from mock', got %v", resp.Output.Content)
	}
}
