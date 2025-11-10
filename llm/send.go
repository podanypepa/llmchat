package llm

import "context"

// Send is a generic function that sends a request to any client implementing the Client interface.
func Send(ctx context.Context, client Client, req *Request) (*Response, error) {
	return client.Send(ctx, req)
}
