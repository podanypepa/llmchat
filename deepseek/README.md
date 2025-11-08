# llmchat/deepseek

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/podanypepa/llmchat/deepseek"
)

func main() {
	c, err := deepseek.NewClient(os.Getenv("DEEPSEEK_API_KEY"))
	if err != nil {
		panic(err)
	}

	req := &deepseek.ChatRequest{
		Model: deepseek.ModelDeepSeekChat,
		Messages: []deepseek.ChatMessage{
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

## Streaming Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/podanypepa/llmchat/deepseek"
)

func main() {
	c, err := deepseek.NewClient(os.Getenv("DEEPSEEK_API_KEY"))
	if err != nil {
		panic(err)
	}

	req := &deepseek.ChatRequest{
		Model: deepseek.ModelDeepSeekChat,
		Messages: []deepseek.ChatMessage{
			{Role: "system", Content: "You are a pirate."},
			{Role: "user", Content: "Tell me a long story about a treasure hunt."},
		},
	}

	ch, err := c.StreamSend(context.TODO(), req)
	if err != nil {
		panic(err)
	}

	for chunk := range ch {
		if len(chunk.Choices) > 0 {
			fmt.Print(chunk.Choices[0].Delta.Content)
		}
	}
	fmt.Println()
}
```
