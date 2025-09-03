package anthropic

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

	req := &Request{
		Model: "",
		Messages: []Message{
			{Role: "system", Content: "You are a pirate."},
			{Role: "user", Content: "Hello!"},
		},
	}

	resp, err := client.Send(context.TODO(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	fmt.Println(resp.Content[0].Text)
}
