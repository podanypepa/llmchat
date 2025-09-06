# llmchat/anthropic

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/podanypepa/llmchat/anthropic"
)

func main() {
	c, err := anthropic.NewClient(os.Getenv("ANTHROPIC_API_KEY"))
	if err != nil {
		panic(err)
	}

	req := &anthropic.ChatRequest{
		Model:     anthropic.ModelClaude35Haiku,
		System:    "You are a helpful assistant.",
		MaxTokens: 300,
		Messages: []anthropic.Message{
			{
				Role:    "user",
				Content: "Write a haiku about the sea.",
			},
		},
	}
	res, err := c.Send(context.TODO(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Content[0].Text)
	fmt.Println("tokens used:", res.Usage.InputTokens+res.Usage.OutputTokens)
}
```
