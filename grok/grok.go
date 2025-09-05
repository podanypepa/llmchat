// Package grok ...
package grok

// ChatCompletionRequest represents a request body for Grok's /chat/completions API.
type ChatCompletionRequest struct {
	Model            string         `json:"model"`                       // e.g. "grok-3", "grok-3-mini", "grok-4", "grok-4-heavy"
	Messages         []ChatMessage  `json:"messages"`                    // conversation history
	MaxTokens        *int           `json:"max_tokens,omitempty"`        // maximum tokens in output
	Temperature      *float64       `json:"temperature,omitempty"`       // randomness (default 1.0)
	TopP             *float64       `json:"top_p,omitempty"`             // nucleus sampling (default 1.0)
	N                *int           `json:"n,omitempty"`                 // number of completions to generate
	Stream           bool           `json:"stream,omitempty"`            // stream response via SSE
	Stop             []string       `json:"stop,omitempty"`              // stop sequences
	PresencePenalty  *float64       `json:"presence_penalty,omitempty"`  // discourage repeats
	FrequencyPenalty *float64       `json:"frequency_penalty,omitempty"` // penalize frequent tokens
	LogitBias        map[string]int `json:"logit_bias,omitempty"`        // bias tokens
	User             string         `json:"user,omitempty"`              // unique user ID for tracking
}

// ChatMessage represents a single conversational turn.
type ChatMessage struct {
	Role    string `json:"role"`    // "system" | "user" | "assistant" | "tool"
	Content string `json:"content"` // plain text content
}

// ChatCompletionResponse represents the response from Grok's /chat/completions API.
type ChatCompletionResponse struct {
	ID      string   `json:"id"`      // e.g. "chatcmpl-12345"
	Object  string   `json:"object"`  // e.g. "chat.completion"
	Created int64    `json:"created"` // unix timestamp
	Model   string   `json:"model"`   // e.g. "grok-3", "grok-4"
	Choices []Choice `json:"choices"` // generated completions
	Usage   *Usage   `json:"usage,omitempty"`
}

// Choice represents one generated completion.
type Choice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`       // final assistant message
	FinishReason string      `json:"finish_reason"` // "stop", "length", "content_filter", "tool_calls"
}

// Usage provides token accounting for the request.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
