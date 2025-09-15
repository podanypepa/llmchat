// Package perplexity provides functions to calculate the perplexity of a language model.
package perplexity

// ChatRequest represents a request body for Perplexity's /chat/completions API.
type ChatRequest struct {
	Model            string         `json:"model"`                         // e.g. "perplexity/sonar", "perplexity/sonar-pro"
	Messages         []ChatMessage  `json:"messages"`                      // conversation history
	MaxTokens        *int           `json:"max_tokens,omitzero"`           // maximum output tokens
	Temperature      *float64       `json:"temperature,omitzero"`          // randomness (default ~1.0)
	TopP             *float64       `json:"top_p,omitzero"`                // nucleus sampling
	N                *int           `json:"n,omitzero"`                    // number of completions
	Stream           bool           `json:"stream,omitempty"`              // SSE streaming
	Stop             []string       `json:"stop,omitempty,omitzero"`       // stop sequences
	PresencePenalty  *float64       `json:"presence_penalty,omitzero"`     // discourage repetition
	FrequencyPenalty *float64       `json:"frequency_penalty,omitzero"`    // penalize frequent tokens
	LogitBias        map[string]int `json:"logit_bias,omitempty,omitzero"` // bias tokens
	User             string         `json:"user,omitempty"`                // unique user identifier
}

// ChatMessage represents a single conversation turn.
type ChatMessage struct {
	Role    string `json:"role"`    // "system" | "user" | "assistant" | "tool"
	Content string `json:"content"` // plain text content
}

// ChatResponse represents the response from Perplexity's /chat/completions API.
type ChatResponse struct {
	ID      string   `json:"id"`      // unique completion ID
	Object  string   `json:"object"`  // usually "chat.completion"
	Created int64    `json:"created"` // unix timestamp (seconds)
	Model   string   `json:"model"`   // e.g. "perplexity/sonar-pro"
	Choices []Choice `json:"choices"`
	Usage   *Usage   `json:"usage,omitempty"`
}

// Choice is one possible generated completion.
type Choice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`       // assistant reply
	FinishReason string      `json:"finish_reason"` // "stop", "length", "content_filter", "tool_calls"
}

// Usage reports token accounting.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
