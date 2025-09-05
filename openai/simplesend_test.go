package chatgpt

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleSend(t *testing.T) {
	client, err := NewClient(os.Getenv("OPENAI_API_KEY"))
	assert.NoError(t, err)
	assert.NotNil(t, client)

	resp, err := client.SimpleSend(context.TODO(), "Hello!")
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	fmt.Println(resp.Choices[0].Message.Content)
	fmt.Println("total tokens:", resp.Usage.TotalTokens)
}

