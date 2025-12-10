// Package main demonstrates how to use the DeepSeek LLM client to send a chat message
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/podanypepa/llmchat/deepseek"
	"github.com/podanypepa/llmchat/llm"
)

func main() {
	client, err := deepseek.NewClient(os.Getenv("DEEPSEEK_API_KEY"))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	req := &llm.Request{
		Model: deepseek.ModelDeepSeekChat,
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
