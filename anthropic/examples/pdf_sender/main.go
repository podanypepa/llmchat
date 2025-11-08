package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/podanypepa/llmchat/anthropic"
	"rsc.io/pdf" // You need to install this dependency: go get rsc.io/pdf
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a path to a PDF file.")
		os.Exit(1)
	}

	pdfPath := os.Args[1]

	// Open and read the PDF file
	file, err := pdf.Open(pdfPath)
	if err != nil {
		panic(fmt.Errorf("failed to open PDF file: %w", err))
	}

	var content strings.Builder
	for i := 1; i <= file.NumPage(); i++ {
		page := file.Page(i)
		if page.V.IsNull() {
			continue
		}
		texts := page.Content().Text
		for _, text := range texts {
			content.WriteString(text.S)
		}
		content.WriteString("\n")
	}

	c, err := anthropic.NewClient(os.Getenv("ANTHROPIC_API_KEY"))
	if err != nil {
		panic(err)
	}

	req := &anthropic.ChatRequest{
		Model:     anthropic.ModelClaude3Opus,
		MaxTokens: 1024,
		Messages: []anthropic.Message{
			{
				Role:    "user",
				Content: fmt.Sprintf("Please summarize the following document:\n\n%s", content.String()),
			},
		},
	}

	res, err := c.Send(context.TODO(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println("Summary:")
	fmt.Println(res.Content[0].Text)
	fmt.Println("\nTokens used:", res.Usage.InputTokens+res.Usage.OutputTokens)
}
