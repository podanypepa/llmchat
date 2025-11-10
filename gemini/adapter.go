package gemini

import (
	"context"

	"github.com/podanypepa/llmchat/llm"
)

// Send implements the llm.Client interface for Gemini.
func (c *Client) Send(ctx context.Context, req *llm.Request) (*llm.Response, error) {
	geminiReq := &ChatRequest{
		Contents: []Content{
			{
				Role:  "user",
				Parts: []Part{{Text: req.Messages[0].Content}},
			},
		},
	}
	// Note: Gemini API has a different structure for messages (Contents).
	// This adapter simplifies it to a single user message for now.
	// For more complex conversations, this mapping would need to be more sophisticated.
	resp, err := c.send(ctx, geminiReq)
	if err != nil {
		return nil, err
	}

	return &llm.Response{
		Content: resp.Candidates[0].Content.Parts[0].Text,
		Metadata: llm.Metadata{
			Usage: llm.Usage{
				PromptTokens:     resp.UsageMetadata.PromptTokenCount,
				CompletionTokens: resp.UsageMetadata.CandidatesTokenCount,
				TotalTokens:      resp.UsageMetadata.TotalTokenCount,
			},
		},
	}, nil
}

// StreamSend implements the llm.Client interface for Gemini.
func (c *Client) StreamSend(ctx context.Context, req *llm.Request) (<-chan llm.StreamChunk, error) {
	geminiReq := &ChatRequest{
		Contents: []Content{
			{
				Role:  "user",
				Parts: []Part{{Text: req.Messages[0].Content}},
			},
		},
	}
	geminiStream, err := c.streamSend(ctx, geminiReq)
	if err != nil {
		return nil, err
	}

	llmStream := make(chan llm.StreamChunk)
	go func() {
		defer close(llmStream)
		for chunk := range geminiStream {
			if len(chunk.Candidates) > 0 && len(chunk.Candidates[0].Content.Parts) > 0 {
				llmStream <- llm.StreamChunk{
					Content: chunk.Candidates[0].Content.Parts[0].Text,
				}
			}
		}
	}()

	return llmStream, nil
}
