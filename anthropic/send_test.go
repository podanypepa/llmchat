package anthropic

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	client, err := NewClient(os.Getenv("ANTHROPIC_API_KEY"))
	assert.NoError(t, err)
	assert.NotNil(t, client)

	req := &Request{
		Model:     ModelClaude35Haiku,
		System:    "You are a helpful assistant.",
		MaxTokens: 300,
		Messages: []Message{
			{
				Role:    "user",
				Content: "Write a haiku about the sea.",
			},
		},
	}

	resp, err := client.Send(context.TODO(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	if err != nil || resp == nil {
		t.FailNow()
	}

	fmt.Println(resp.Content[0].Text)
	fmt.Println("input tokens:", resp.Usage.InputTokens)
	fmt.Println("output tokens:", resp.Usage.OutputTokens)
	fmt.Println("total tokens:", resp.Usage.InputTokens+resp.Usage.OutputTokens)
}
