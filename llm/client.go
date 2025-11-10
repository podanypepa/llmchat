// Package llm defines the interface for LLM clients.
package llm

import "context"

// Client is the interface that all LLM providers must implement.
type Client interface {
	Send(ctx context.Context, req *Request) (*Response, error)
	StreamSend(ctx context.Context, req *Request) (<-chan StreamChunk, error)
}
