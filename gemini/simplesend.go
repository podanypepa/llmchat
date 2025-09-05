package gemini

import "context"

// SimpleSend is a helper function that sends a simple message to the ChatGPT API
func (c *Client) SimpleSend(ctx context.Context, message string) (*GenerateContentResponse, error) {
	req := &GenerateContentRequest{
		Contents: []Content{
			{
				Role: "user",
				Parts: []Part{
					{Text: message},
				},
			},
		},
		GenerationConfig: &GenerationConfig{
			// TODO: must be in client config or req
			Temperature: 0.7,
		},
	}

	return c.Send(ctx, req)
}

// SimpleSendWithSystem is a helper function that sends a message with a system prompt to the ChatGPT API
func (c *Client) SimpleSendWithSystem(ctx context.Context, system, message string) (*GenerateContentResponse, error) {
	req := &GenerateContentRequest{
		Contents: []Content{
			{
				Role: "system",
				Parts: []Part{
					{Text: system},
				},
			},
			{
				Role: "user",
				Parts: []Part{
					{Text: message},
				},
			},
		},
		GenerationConfig: &GenerationConfig{
			// TODO: must be in client config or req
			Temperature: 0.7,
		},
	}
	return c.Send(ctx, req)
}

// SetModel sets the default model for the client
func (c *Client) SetModel(model string) {
	c.config.DefaultModel = model
}
