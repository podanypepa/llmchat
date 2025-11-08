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
	fmt.Println("tokens used:", res.UsageMetadata.TotalTokenCount)
}
```

## Streaming Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/podanypepa/llmchat/gemini"
)

func main() {
	c, err := gemini.NewClient(os.Getenv("GEMINI_API_KEY"), gemini.GeminiPro)
	if err != nil {
		panic(err)
	}

	req := &gemini.ChatRequest{
		Contents: []gemini.Content{
			{
				Role: "user",
				Parts: []gemini.Part{
					{Text: "Tell me a long story about a robot who discovers music."},
				},
			},
		},
	}

	ch, err := c.StreamSend(context.TODO(), req)
	if err != nil {
		panic(err)
	}

	for resp := range ch {
		text, err := resp.ExtractText()
		if err != nil {
			fmt.Printf("Error extracting text: %v\n", err)
			continue
		}
		fmt.Print(text)
	}
	fmt.Println()
}
```
