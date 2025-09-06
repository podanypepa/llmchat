# llmchat/gemini

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/podanypepa/llmchat/gemini"
)

func main() {
	c, err := gemini.NewClient(os.Getenv("GEMINI_API_KEY"), gemini.GeminI2_5Pro)
	if err != nil {
		panic(err)
	}

	res, err := c.SimpleSend(context.TODO(), "Hello World!")
	if err != nil {
		panic(err)
	}

	fmt.Println(res.ExtractText())

	res, err = c.Send(context.TODO(), &gemini.ChatRequest{
		SystemInstruction: &gemini.Content{
			Role: "user",
			Parts: []gemini.Part{
				{Text: "You are hacker."},
			},
		},
		Contents: []gemini.Content{
			{
				Role: "user",
				Parts: []gemini.Part{
					{Text: "Write a python script that prints 'Hello World!'"},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res.ExtractText())
	fmt.Println("tokens used:", res.UsageMetadata.TotalTokenCount)
}
```
