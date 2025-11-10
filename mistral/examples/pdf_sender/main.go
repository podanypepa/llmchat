package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/podanypepa/llmchat/llm"
	"github.com/podanypepa/llmchat/mistral"
)

func main() {
	client, err := mistral.NewClient(os.Getenv("MISTRAL_API_KEY"))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	req := &llm.Request{
		Model: mistral.ModelMistralSmallLatest,
		Messages: []llm.ChatMessage{
			{
				Role:    "user",
				Content: "Write a short story about a brave knight.",
			},
		},
	}

	resp, err := client.Send(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to send message: %v", err)
	}

	fmt.Println("Response:", resp.Content)
	fmt.Printf("Usage: %+v\n", resp.Metadata.Usage)
}
