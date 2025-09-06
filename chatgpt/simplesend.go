package chatgpt

import "context"

// SimpleSend is a helper function that sends a simple message to the ChatGPT API
func (c *Client) SimpleSend(ctx context.Context, message string) (*ChatResponse, error) {
	req := &ChatRequest{
		Model: Gpt4OMini,
		Messages: []ChatMessage{
			{
				Role:    RoleUser,
				Content: message,
			},
		},
	}

	return c.Send(ctx, req)
}

// SimpleSendWithSystem is a helper function that sends a message with a system prompt to the ChatGPT API
func (c *Client) SimpleSendWithSystem(ctx context.Context, system, message string) (*ChatResponse, error) {
	req := &ChatRequest{
		Model: c.config.DefaultModel,
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

// SetModel sets the default model for the client
func (c *Client) SetModel(model string) {
	c.config.DefaultModel = model
}
