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

## Streaming Example

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
			{Role: "user", Content: "Tell me a long story about a space explorer."},
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

## Image Input Example

```go
package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/podanypepa/llmchat/grok"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a path to an image file.")
		os.Exit(1)
	}
	imagePath := os.Args[1]

	// Read image file
	imageData, err := os.ReadFile(imagePath)
	if err != nil {
		panic(fmt.Errorf("failed to read image file: %w", err))
	}

	// Encode to base64
	base64Image := base64.StdEncoding.EncodeToString(imageData)

	// Detect MIME type
	mimeType := http.DetectContentType(imageData)

	c, err := grok.NewClient(os.Getenv("XAI_API_KEY"))
	if err != nil {
		panic(err)
	}

	req := &grok.ChatRequest{
		Model: grok.ModelGrok4, // Or another model that supports vision
		Messages: []grok.ChatMessage{
			{
				Role: "user",
				Content: []grok.ContentPart{
					{
						Type: "text",
						Text: "What's in this image?",
					},
					{
						Type: "image_url",
						ImageURL: &grok.ImageURL{
							URL: fmt.Sprintf("data:%s;base64,%s", mimeType, base64Image),
						},
					},
				},
			},
		},
		MaxTokens: grok.Ptr(300),
	}

	res, err := c.Send(context.TODO(), req)
	if err != nil {
		panic(err)
	}

	// Note: The response 'Content' is a string, so we need to cast it.
	if content, ok := res.Choices[0].Message.Content.(string); ok {
		fmt.Println(content)
	}
	fmt.Println("tokens used:", res.Usage.TotalTokens)
}
```
