# llmchat

**One Go package for all major LLM chat providers** — Anthropic, OpenAI/ChatGPT, Gemini, Mistral, Perplexity, DeepSeek, and Grok.

---

## 🚀 Why llmchat?

- Unified interface to interact with multiple LLM providers.
- Simple and consistent usage across different APIs.
- Easily extendable: adding or swapping providers is plug-and-play.
- Written in Go, ideal for backend services, bots, and CLI tools.

---

## ⚡ Streaming Support

All supported LLM providers (Anthropic, OpenAI/ChatGPT, Gemini, Mistral, Perplexity, DeepSeek, and Grok) now offer streaming functionality, allowing for real-time, incremental responses.

---

## 📚 Examples

Each provider now includes an example application (`examples/pdf_sender`) that demonstrates how to analyze the content of a PDF file using its respective LLM model. This allows you to easily integrate PDF summarization or analysis into your Go applications.

---

## 👨‍💻 Quick Start

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

	res, err = c.Send(context.TODO(), &gemini.ChatRequest{
		SystemInstruction: &gemini.Content{
			Role: "user",
			Parts: []gemini.Part{
				{Text: "You are hacker."},
			},
		},
		Contents: []gemini.Content{
			{
				Role: "user",
				Parts: []gemini.Part{
					{Text: "Write a python script that prints 'Hello World!'"},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(res.ExtractText())
    fmt.Println("tokens used:", res.UsageMetadata.TotalTokenCount)
}
```
