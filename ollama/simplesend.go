package ollama

import "context"

const (
	// RoleSystem is the system role for chat messages.
	RoleSystem = "system"
	// RoleUser is the user role for chat messages.
	RoleUser = "user"
	// RoleAssistant is the assistant role for chat messages.
	RoleAssistant = "assistant"
)

// SimpleSend is a helper function that sends a simple message to the Ollama API
func (c *Client) SimpleSend(ctx context.Context, model, message string) (*ChatResponse, error) {
	req := &ChatRequest{
		Model: model,
		Messages: []ChatMessage{
			{
				Role:    RoleUser,
				Content: message,
			},
		},
	}

	return c.Send(ctx, req)
}

// SimpleSendWithSystem is a helper function that sends a message with a system prompt to the Ollama API
func (c *Client) SimpleSendWithSystem(ctx context.Context, model, system, message string) (*ChatResponse, error) {
	req := &ChatRequest{
		Model: model,
		Messages: []ChatMessage{
			{
				Role:    RoleSystem,
				Content: system,
			},
			{
				Role:    RoleUser,
				Content: message,
			},
		},
	}

	return c.Send(ctx, req)
}