package mistral

import "context"

// SimpleSend is a helper function that sends a simple message to the ChatGPT API
func (c *Client) SimpleSend(ctx context.Context, message string) (*ChatCompletionResponse, error) {
	req := &ChatRequest{
		Model: DefaultModel, // TODO: must be in client config or req
		Messages: []ChatMessage{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: message},
		},
	}

	return c.Send(ctx, req)
}

// SimpleSendWithSystem is a helper function that sends a message with a system prompt to the ChatGPT API
func (c *Client) SimpleSendWithSystem(ctx context.Context, system, message string) (*ChatCompletionResponse, error) {
	req := &ChatRequest{
		Model: DefaultModel, // TODO: must be in client config or req
		Messages: []ChatMessage{
			{Role: "system", Content: system},
			{Role: "user", Content: message},
		},
	}

	return c.Send(ctx, req)
}

// SetModel sets the default model for the client
func (c *Client) SetModel(model string) {
	c.config.DefaultModel = model
}
