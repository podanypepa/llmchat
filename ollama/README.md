# llmchat/ollama

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/podanypepa/llmchat/ollama"
)

func main() {
	client, err := ollama.NewClient()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	model := "llama3" // Make sure you have this model installed in Ollama
	prompt := "Why is the sky blue?"

	req := &ollama.ChatRequest{
		Model: model,
		Messages: []ollama.ChatMessage{
			{
				Role:    ollama.RoleUser,
				Content: prompt,
			},
		},
	}

	fmt.Printf("Sending prompt to model %s: '%s'\n", model, prompt)

	res, err := client.Send(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to send message: %v", err)
	}

	fmt.Println("Response:")
	fmt.Println(res.Message.Content)
}
```

## Streaming Example

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/podanypepa/llmchat/ollama"
)

func main() {
	client, err := ollama.NewClient()
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	model := "llama3" // Make sure you have this model installed in Ollama
	prompt := "Tell me a long story about a treasure hunt."

	req := &ollama.ChatRequest{
		Model: model,
		Messages: []ollama.ChatMessage{
			{
				Role:    ollama.RoleUser,
				Content: prompt,
			},
		},
	}

	fmt.Printf("Sending prompt to model %s: '%s'\n", model, prompt)

	ch, err := client.StreamSend(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to send message: %v", err)
	}

	fmt.Println("Response:")
	for chunk := range ch {
		fmt.Print(chunk.Message.Content)
	}
	fmt.Println()
}
```
