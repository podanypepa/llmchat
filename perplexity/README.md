# llmchat/perplexity

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/podanypepa/llmchat/perplexity"
)

func main() {
	c, err := perplexity.NewClient(os.Getenv("PERPLEXITY_API_KEY"))
	if err != nil {
		panic(err)
	}

	req := &perplexity.ChatRequest{
		Model: perplexity.ModelPerplexitySonarPro,
		Messages: []perplexity.ChatMessage{
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
