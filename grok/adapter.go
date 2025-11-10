package grok

import (
	"context"

	"github.com/podanypepa/llmchat/llm"
)

// Send implements the llm.Client interface for Grok.
func (c *Client) Send(ctx context.Context, req *llm.Request) (*llm.Response, error) {
	grokReq := &ChatRequest{
		Model:    req.Model,
		Messages: make([]ChatMessage, len(req.Messages)),
	}
	for i, msg := range req.Messages {
		grokReq.Messages[i] = ChatMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	resp, err := c.send(ctx, grokReq)
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

// StreamSend implements the llm.Client interface for Grok.
func (c *Client) StreamSend(ctx context.Context, req *llm.Request) (<-chan llm.StreamChunk, error) {
	grokReq := &ChatRequest{
		Model:    req.Model,
		Messages: make([]ChatMessage, len(req.Messages)),
	}
	for i, msg := range req.Messages {
		grokReq.Messages[i] = ChatMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	grokStream, err := c.streamSend(ctx, grokReq)
	if err != nil {
		return nil, err
	}

	llmStream := make(chan llm.StreamChunk)
	go func() {
		defer close(llmStream)
		for chunk := range grokStream {
			if len(chunk.Choices) > 0 {
				llmStream <- llm.StreamChunk{
					Content: chunk.Choices[0].Delta.Content,
				}
			}
		}
	}()

	return llmStream, nil
}