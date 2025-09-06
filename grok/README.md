# llmchat/grok

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/podanypepa/llmchat/grok"
)

func main() {
	c, err := grok.NewClient(os.Getenv("XAI_API_KEY"))
	if err != nil {
		panic(err)
	}

	req := &grok.ChatRequest{
		Model: grok.ModelGrok4,
		Messages: []grok.ChatMessage{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: "Write a haiku about Go concurrency."},
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
