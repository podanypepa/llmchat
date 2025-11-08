// Package mistral provides a client for interacting with the Mistral workflow service.
package mistral

// ChatRequest represents a request body for Mistral's /chat/completions API.
type ChatRequest struct {
	Model            string         `json:"model"`                       // e.g. "mistral-7b-instruct", "mixtral-8x7b-instruct", "mistral-small-3.1"
	Messages         []ChatMessage  `json:"messages"`                    // conversation history
	MaxTokens        *int           `json:"max_tokens,omitempty"`        // maximum tokens in output
	Temperature      *float64       `json:"temperature,omitempty"`       // randomness (default 1.0)
	TopP             *float64       `json:"top_p,omitempty"`             // nucleus sampling (default 1.0)
	N                *int           `json:"n,omitempty"`                 // number of completions to generate
	Stream           bool           `json:"stream,omitempty"`            // stream response via SSE
	Stop             []string       `json:"stop,omitempty"`              // stop sequences
	PresencePenalty  *float64       `json:"presence_penalty,omitempty"`  // discourage repetition
	FrequencyPenalty *float64       `json:"frequency_penalty,omitempty"` // penalize frequent tokens
	LogitBias        map[string]int `json:"logit_bias,omitempty"`        // bias tokens
	User             string         `json:"user,omitempty"`              // user identifier
}

// ChatMessage represents a single conversational turn.
type ChatMessage struct {
	Role    string      `json:"role"`    // "system" | "user" | "assistant" | "tool"
	Content interface{} `json:"content"` // text content
}

// ContentPart represents a part of a multimodal content message.
type ContentPart struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	ImageURL *ImageURL `json:"image_url,omitempty"`
}

// ImageURL represents the URL of an image in a content part.
type ImageURL struct {
	URL    string `json:"url"`
	Detail string `json:"detail,omitempty"`
}

// ChatResponse represents the response from Mistral's /chat/completions API.
type ChatResponse struct {
	ID      string   `json:"id"`      // unique identifier for the completion
	Object  string   `json:"object"`  // usually "chat.completion"
	Created int64    `json:"created"` // unix timestamp
	Model   string   `json:"model"`   // e.g. "mistral-small-3.1"
	Choices []Choice `json:"choices"`
	Usage   *Usage   `json:"usage,omitempty"`
}

// Choice represents one generated completion option.
type Choice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`       // the assistant's reply
	FinishReason string      `json:"finish_reason"` // "stop", "length", "content_filter", "tool_calls"
}

// Usage provides token usage statistics.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// StreamCompletionChunk is a single chunk of a streamed completion.
type StreamCompletionChunk struct {
	ID      string                   `json:"id"`
	Object  string                   `json:"object"`
	Created int64                    `json:"created"`
	Model   string                   `json:"model"`
	Choices []StreamCompletionChoice `json:"choices"`
}

// StreamCompletionChoice is a single choice in a streamed completion.
type StreamCompletionChoice struct {
	Index        int                   `json:"index"`
	Delta        StreamCompletionDelta `json:"delta"`
	FinishReason string                `json:"finish_reason"`
}

// StreamCompletionDelta is the delta of a streamed completion.
type StreamCompletionDelta struct {
	Content string `json:"content"`
}

