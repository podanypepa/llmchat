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

## Streaming Example

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
				Content: "Write a long story about a robot who discovers music.",
			},
		},
	}

	ch, err := c.StreamSend(context.TODO(), req)
	if err != nil {
		panic(err)
	}

	for event := range ch {
		if event.Type == "content_block_delta" && event.Delta.Type == "text_delta" {
			fmt.Print(event.Delta.Text)
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

	"github.com/podanypepa/llmchat/anthropic"
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

	c, err := anthropic.NewClient(os.Getenv("ANTHROPIC_API_KEY"))
	if err != nil {
		panic(err)
	}

	req := &anthropic.ChatRequest{
		Model:     anthropic.ModelClaude3Opus, // Opus is recommended for image inputs
		MaxTokens: 300,
		Messages: []anthropic.Message{
			{
				Role: "user",
				Content: []anthropic.ContentBlock{
					{
						Type: "text",
						Text: "What's in this image?",
					},
					{
						Type: "image",
						Source: &anthropic.ImageSource{
							Type:      "base64",
							MediaType: &mimeType,
							Data:      &base64Image,
						},
					},
				},
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
