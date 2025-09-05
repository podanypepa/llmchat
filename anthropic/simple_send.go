package anthropic

import "context"

// SimpleSend is a helper function that sends a simple message to the ChatGPT API
func (c *Client) SimpleSend(ctx context.Context, message string) (*Response, error) {
	req := &Request{
		Model: "",
		Messages: []Message{
			{
				Role:    RoleUser,
				Content: []string{message},
			},
		},
	}

	return c.Send(ctx, req)
}

// SimpleSendWithSystem is a helper function that sends a message with a system prompt to the ChatGPT API
func (c *Client) SimpleSendWithSystem(ctx context.Context, system, message string) (*Response, error) {
	req := &Request{
		Model: "",
		Messages: []Message{
			{
				Role:    RoleSystem,
				Content: []string{system},
			},
			{
				Role:    RoleUser,
				Content: []string{message},
			},
		},
	}

	return c.Send(ctx, req)
}
