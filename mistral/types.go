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
	Role    string `json:"role"`    // "system" | "user" | "assistant" | "tool"
	Content string `json:"content"` // text content
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
