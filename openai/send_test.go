package chatgpt

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	client, err := NewClient(os.Getenv("OPENAI_API_KEY"))
	assert.NoError(t, err)
	assert.NotNil(t, client)

	req := &ChatRequest{
		Model: Gpt4OMini,
		Messages: []ChatMessage{
			{Role: "system", Content: "You are a pirate."},
			{Role: "user", Content: "Hello!"},
		},
	}

	resp, err := client.Send(context.TODO(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	fmt.Println(resp.Choices[0].Message.Content)
	fmt.Println("total tokens:", resp.Usage.TotalTokens)
}
