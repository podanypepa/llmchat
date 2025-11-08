package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/podanypepa/llmchat/mistral"
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

	c, err := mistral.NewClient(os.Getenv("MISTRAL_API_KEY"))
	if err != nil {
		panic(err)
	}

	req := &mistral.ChatRequest{
		Model: mistral.ModelMistralLargeLatest,
		Messages: []mistral.ChatMessage{
			{
				Role:    "system",
				Content: "You are an assistant that summarizes PDF documents.",
			},
			{
				Role:    "user",
				Content: fmt.Sprintf("Please summarize the following document:\n\n%s", content.String()),
			},
		},
		MaxTokens: mistral.Ptr(1024),
	}

	res, err := c.Send(context.TODO(), req)
	if err != nil {
		panic(err)
	}

	if responseContent, ok := res.Choices[0].Message.Content.(string); ok {
		fmt.Println("Summary:")
		fmt.Println(responseContent)
	}
	fmt.Println("\nTokens used:", res.Usage.TotalTokens)
}
