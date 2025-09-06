# llmchat/mistral

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/podanypepa/llmchat/mistral"
)

func main() {
	c, err := mistral.NewClient(os.Getenv("MISTRAL_API_KEY"))
	if err != nil {
		panic(err)
	}

	req := &mistral.ChatRequest{
		Model: mistral.ModelMistralSmallLatest,
		Messages: []mistral.ChatMessage{
			{Role: "system", Content: "You are a pirate."},
			{Role: "user", Content: "Hello!"},
		},
	}

	res, err := c.Send(context.TODO(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Choices[0].Message.Content)
	fmt.Println("tokens used:", res.Usage.TotalTokens)
}
```
