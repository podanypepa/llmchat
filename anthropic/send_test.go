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
		Model: "",
		Messages: []Message{
			{Role: "system", Content: []string{"You are a helpful assistant."}},
			{Role: "user", Content: []string{"Write a haiku about the sea."}},
		},
	}

	resp, err := client.Send(context.TODO(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	if err != nil || resp == nil {
		t.FailNow()
	}

	fmt.Println(resp.Content[0].Text)
}
