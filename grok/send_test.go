package grok

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	client, err := NewClient(os.Getenv("XAI_API_KEY"))
	assert.NoError(t, err)
	assert.NotNil(t, client)

	req := &ChatRequest{
		Model: DefaultModel,
		Messages: []ChatMessage{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: "Write a haiku about Go concurrency."},
		},
	}
	resp, err := client.Send(context.TODO(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	fmt.Println(resp.Choices[0].Message.Content)
	fmt.Println("total tokens:", resp.Usage.TotalTokens)
}
