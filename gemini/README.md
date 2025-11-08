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
	c, err := gemini.NewClient(os.Getenv("GEMINI_API_KEY"), gemini.GeminiPro)
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

## Image Input Example

```go
package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/podanypepa/llmchat/gemini"
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

	c, err := gemini.NewClient(os.Getenv("GEMINI_API_KEY"), gemini.GeminiPro) // Ensure you use a model that supports vision
	if err != nil {
		panic(err)
	}

	req := &gemini.ChatRequest{
		Contents: []gemini.Content{
			{
				Role: "user",
				Parts: []gemini.Part{
					{
						Text: "What's in this image?",
					},
					{
						InlineData: &gemini.InlineData{
							MimeType: mimeType,
							Data:     base64Image,
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

	text, err := res.ExtractText()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
	fmt.Println("tokens used:", res.UsageMetadata.TotalTokenCount)
}
```
