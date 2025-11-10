package anthropic

import (
	"context"
	"fmt"

	"github.com/podanypepa/llmchat/llm"
)

// Send implements the llm.Client interface for Anthropic.
func (c *Client) Send(ctx context.Context, req *llm.Request) (*llm.Response, error) {
	anthropicReq, err := toAnthropicRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to anthropic request: %w", err)
	}

	resp, err := c.send(ctx, anthropicReq)
	if err != nil {
		return nil, err
	}

	return toLLMResponse(resp)
}

// StreamSend implements the llm.Client interface for Anthropic.
func (c *Client) StreamSend(ctx context.Context, req *llm.Request) (<-chan llm.StreamChunk, error) {
	anthropicReq, err := toAnthropicRequest(req)
	if err != nil {
		return nil, fmt.Errorf("failed to convert to anthropic request: %w", err)
	}

	anthropicStream, err := c.streamSend(ctx, anthropicReq)
	if err != nil {
		return nil, err
	}

	llmStream := make(chan llm.StreamChunk)
	go func() {
		defer close(llmStream)
		for chunk := range anthropicStream {
			llmStream <- llm.StreamChunk{
				Content: chunk.Delta.Text,
			}
		}
	}()

	return llmStream, nil
}

func toAnthropicRequest(req *llm.Request) (*ChatRequest, error) {
	messages := make([]Message, len(req.Messages))
	for i, msg := range req.Messages {
		messages[i] = Message{
			Role:    Role(msg.Role),
			Content: msg.Content,
		}
	}

	return &ChatRequest{
		Model:     req.Model,
		Messages:  messages,
		MaxTokens: 4096, // Default value, can be overridden
	}, nil
}

func toLLMResponse(resp *ChatResponse) (*llm.Response, error) {
	if len(resp.Content) == 0 {
		return nil, fmt.Errorf("no content in response")
	}
	return &llm.Response{
		Content: resp.Content[0].Text,
		Metadata: llm.Metadata{
			Usage: llm.Usage{
				PromptTokens:     resp.Usage.InputTokens,
				CompletionTokens: resp.Usage.OutputTokens,
			},
		},
	}, nil
}
