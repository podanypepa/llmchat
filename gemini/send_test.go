package gemini

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSend(t *testing.T) {
	client, err := NewClient(os.Getenv("GEMINI_API_KEY"), GeminI2_5Flash)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	req := &ChatRequest{
		Contents: []Content{
			{
				Role: "user",
				Parts: []Part{
					{Text: "Write a haiku about Go programming."},
				},
			},
		},
		GenerationConfig: &GenerationConfig{
			Temperature: 0.7,
		},
	}

	resp, err := client.Send(context.TODO(), req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	fmt.Println(resp.ExtractText())
}
