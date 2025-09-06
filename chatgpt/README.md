# llmchat/chatgpt

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/podanypepa/llmchat/chatgpt"
)

func main() {
	c, err := chatgpt.NewClient(os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		panic(err)
	}

	req := &chatgpt.ChatRequest{
		Model: chatgpt.Gpt4OMini,
		Messages: []chatgpt.ChatMessage{
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
