package perplexity

import (
	"context"

	"github.com/podanypepa/llmchat/llm"
)

// Send implements the llm.Client interface for Perplexity.
func (c *Client) Send(ctx context.Context, req *llm.Request) (*llm.Response, error) {
	perplexityReq := &ChatRequest{
		Model:    req.Model,
		Messages: make([]ChatMessage, len(req.Messages)),
	}
	for i, msg := range req.Messages {
		perplexityReq.Messages[i] = ChatMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	resp, err := c.send(ctx, perplexityReq)
	if err != nil {
		return nil, err
	}

	return &llm.Response{
		Content: resp.Choices[0].Message.Content.(string),
		Metadata: llm.Metadata{
			Usage: llm.Usage{
				PromptTokens:     resp.Usage.PromptTokens,
				CompletionTokens: resp.Usage.CompletionTokens,
				TotalTokens:      resp.Usage.TotalTokens,
			},
		},
	}, nil
}

// StreamSend implements the llm.Client interface for Perplexity.
func (c *Client) StreamSend(ctx context.Context, req *llm.Request) (<-chan llm.StreamChunk, error) {
	perplexityReq := &ChatRequest{
		Model:    req.Model,
		Messages: make([]ChatMessage, len(req.Messages)),
	}
	for i, msg := range req.Messages {
		perplexityReq.Messages[i] = ChatMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	perplexityStream, err := c.streamSend(ctx, perplexityReq)
	if err != nil {
		return nil, err
	}

	llmStream := make(chan llm.StreamChunk)
	go func() {
		defer close(llmStream)
		for chunk := range perplexityStream {
			if len(chunk.Choices) > 0 {
				llmStream <- llm.StreamChunk{
					Content: chunk.Choices[0].Delta.Content,
				}
			}
		}
	}()

	return llmStream, nil
}