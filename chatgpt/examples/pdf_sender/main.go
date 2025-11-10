// Package main implements a command-line tool that reads a PDF file,
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/podanypepa/llmchat/chatgpt"
	"github.com/podanypepa/llmchat/llm"
	"github.com/podanypepa/llmchat/pkg/pdftools"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a path to a PDF file.")
	}
	pdfPath := os.Args[1]

	text, err := pdftools.ReadPdf(pdfPath)
	if err != nil {
		log.Fatalf("failed to read PDF: %v", err)
	}

	c, err := chatgpt.NewClient(os.Getenv("OPENAI_API_KEY"))
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}

	req := &llm.Request{
		Model: "gpt-4o",
		Messages: []llm.ChatMessage{
			{
				Role:    "user",
				Content: fmt.Sprintf("what is in the pdf file? '%s'", string(text)),
			},
		},
	}

	res, err := llm.Send(context.TODO(), c, req)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.Content)
	fmt.Println("tokens used:", res.Metadata.Usage.TotalTokens)
}
