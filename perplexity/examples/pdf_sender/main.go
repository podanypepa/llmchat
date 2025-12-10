// Package main demonstrates how to use the Perplexity LLM client
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/podanypepa/llmchat/llm"
	"github.com/podanypepa/llmchat/perplexity"
)

func main() {
	client, err := perplexity.NewClient(os.Getenv("PERPLEXITY_API_KEY"))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	req := &llm.Request{
		Model: "llama-3-sonar-small-32k-online",
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
