package anthropic

import "context"

// SimpleSend is a helper function that sends a simple message to the ChatGPT API
func (c *Client) SimpleSend(ctx context.Context, message string) (*Response, error) {
	req := &Request{
		Model: "",
		Messages: []Message{
			{
				Role:    RoleUser,
				Content: message,
			},
		},
	}

	return c.Send(ctx, req)
}
