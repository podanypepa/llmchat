package gemini

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleSend(t *testing.T) {
	t.Skip("Skipping integration test")
	client, err := NewClient(os.Getenv("GEMINI_API_KEY"), GeminiPro)
	assert.NoError(t, err)
	assert.NotNil(t, client)

	resp, err := client.SimpleSend(context.TODO(), "Hello!")
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	fmt.Println(resp.Candidates[0].Content.Parts[0].Text)
	fmt.Println("total tokens", resp.UsageMetadata.TotalTokenCount)
}
