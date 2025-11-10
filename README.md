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

This repository includes various example applications to help you get started with `llmchat`:

*   **PDF Analysis Examples**: Each provider now includes an example application (`examples/pdf_sender`) that demonstrates how to analyze the content of a PDF file using its respective LLM model. This allows you to easily integrate PDF summarization or analysis into your Go applications.

*   **Image Generation Examples**: New examples have been added for providers that support image generation, demonstrating how to send a prompt to the API, process the response, and save the generated image to disk.

**Providers supporting image generation:**
*   **ChatGPT (OpenAI)**: Yes, using models like DALL-E 3.
*   **DeepSeek**: Yes, using models like Janus-Pro-7B.
*   **Gemini (Google)**: Yes, using models like Imagen-4.
*   **Grok (xAI)**: Yes, using models like grok-2-image-1212.
*   **Mistral AI**: Yes, using multimodal models like Pixtral.

**Providers NOT supporting native image generation:**
*   **Anthropic**: Does not support native image generation. Its models are multimodal for image *analysis*, not creation.
*   **Perplexity AI**: Does not offer direct image generation through its API. It integrates other models for image generation on its platform, but not via its public API.

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

	res, err := c.Send(context.TODO(), &gemini.ChatRequest{
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
