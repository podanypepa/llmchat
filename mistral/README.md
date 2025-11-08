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

## Streaming Example

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
			{Role: "user", Content: "Tell me a long story about a sea adventure."},
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

	"github.com/podanypepa/llmchat/mistral"
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

	c, err := mistral.NewClient(os.Getenv("MISTRAL_API_KEY"))
	if err != nil {
		panic(err)
	}

	req := &mistral.ChatRequest{
		Model: mistral.ModelMistralLargeLatest, // Or another model that supports vision
		Messages: []mistral.ChatMessage{
			{
				Role: "user",
				Content: []mistral.ContentPart{
					{
						Type: "text",
						Text: "What's in this image?",
					},
					{
						Type: "image_url",
						ImageURL: &mistral.ImageURL{
							URL: fmt.Sprintf("data:%s;base64,%s", mimeType, base64Image),
						},
					},
				},
			},
		},
		MaxTokens: mistral.Ptr(300),
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
