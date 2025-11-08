// Package main implements a command-line tool that reads a PDF file,
package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/podanypepa/llmchat/gemini"
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

	c, err := gemini.NewClient(os.Getenv("GEMINI_API_KEY"), gemini.GeminiPro)
	if err != nil {
		panic(err)
	}

	req := &gemini.ChatRequest{
		Contents: []gemini.Content{
			{
				Role: "user",
				Parts: []gemini.Part{
					{Text: fmt.Sprintf("Please summarize the following document:\n\n%s", content.String())},
				},
			},
		},
	}

	res, err := c.Send(context.TODO(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println("Summary:")
	text, err := res.ExtractText()
	if err != nil {
		panic(err)
	}
	fmt.Println(text)
	fmt.Println("\nTokens used:", res.UsageMetadata.TotalTokenCount)
}
