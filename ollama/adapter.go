package ollama

import (
	"context"

	"github.com/podanypepa/llmchat/llm"
)

// Send implements the llm.Client interface for Ollama.
func (c *Client) Send(ctx context.Context, req *llm.Request) (*llm.Response, error) {
	ollamaReq := &ChatRequest{
		Model:    req.Model,
		Messages: make([]ChatMessage, len(req.Messages)),
	}
	for i, msg := range req.Messages {
		ollamaReq.Messages[i] = ChatMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	resp, err := c.send(ctx, ollamaReq)
	if err != nil {
		return nil, err
	}

	return &llm.Response{
		Content: resp.Message.Content,
		Metadata: llm.Metadata{
			Usage: llm.Usage{
				PromptTokens:     resp.PromptEvalCount,
				CompletionTokens: resp.EvalCount,
				TotalTokens:      resp.PromptEvalCount + resp.EvalCount,
			},
		},
	}, nil
}

// StreamSend implements the llm.Client interface for Ollama.
func (c *Client) StreamSend(ctx context.Context, req *llm.Request) (<-chan llm.StreamChunk, error) {
	ollamaReq := &ChatRequest{
		Model:    req.Model,
		Messages: make([]ChatMessage, len(req.Messages)),
	}
	for i, msg := range req.Messages {
		ollamaReq.Messages[i] = ChatMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	ollamaStream, err := c.streamSend(ctx, ollamaReq)
	if err != nil {
		return nil, err
	}

	llmStream := make(chan llm.StreamChunk)
	go func() {
		defer close(llmStream)
		for chunk := range ollamaStream {
			llmStream <- llm.StreamChunk{
				Content: chunk.Message.Content,
			}
		}
	}()

	return llmStream, nil
}