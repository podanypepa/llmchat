// Package deepseek provides utilities for deep searching within nested data structures.
package deepseek

// ChatRequest represents the request body for DeepSeek's /chat/completions endpoint.
type ChatRequest struct {
	Model            string             `json:"model"`                 // e.g. "deepseek-chat", "deepseek-v3.1"
	Messages         []ChatMessage      `json:"messages"`              // conversation history
	MaxTokens        *int               `json:"max_tokens,omitempty"`  // optional cap on output tokens
	Temperature      *float64           `json:"temperature,omitempty"` // default ~1.0
	TopP             *float64           `json:"top_p,omitempty"`       // nucleus sampling
	Stream           bool               `json:"stream,omitempty"`      // stream responses via SSE
	Stop             []string           `json:"stop,omitempty"`        // stop sequences
	PresencePenalty  *float64           `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64           `json:"frequency_penalty,omitempty"`
	LogitBias        map[string]float64 `json:"logit_bias,omitempty"`
	User             string             `json:"user,omitempty"` // opaque user ID for tracking
}

// ChatMessage represents one role/content message in the conversation.
type ChatMessage struct {
	Role    string `json:"role"`    // "system" | "user" | "assistant" | "tool"
	Content string `json:"content"` // message text
}

// ReasoningRequest represents a request body for DeepSeek reasoning models (R1 family).
// It extends the chat completion request with reasoning-specific options.
type ReasoningRequest struct {
	Model       string        `json:"model"`                 // e.g. "deepseek-r1", "deepseek-r1-zero"
	Messages    []ChatMessage `json:"messages"`              // conversation context
	MaxTokens   *int          `json:"max_tokens,omitempty"`  // cap on output tokens
	Temperature *float64      `json:"temperature,omitempty"` // sampling temperature
	TopP        *float64      `json:"top_p,omitempty"`       // nucleus sampling
	Stream      bool          `json:"stream,omitempty"`      // enable SSE streaming
	Stop        []string      `json:"stop,omitempty"`        // stop sequences
	User        string        `json:"user,omitempty"`        // external user id for tracking
	// Reasoning-specific options:
	Reasoning ReasoningConfig `json:"reasoning"`
}

// ReasoningConfig allows controlling reasoning mode and thinking traces.
type ReasoningConfig struct {
	// EnableThinking tells the model to generate internal "thinking" tokens
	// (they can be optionally returned in the response if allowed).
	EnableThinking bool `json:"enable_thinking"`
	// ThinkingBudget sets the max number of tokens allocated for reasoning traces.
	// Must be <= MaxTokens. Common values: 1024–4096.
	ThinkingBudget int `json:"thinking_budget,omitempty"`
	// ReturnThinking controls whether reasoning traces are included in the API response.
	// Default is false (final answer only).
	ReturnThinking bool `json:"return_thinking,omitempty"`
}

// ChatResponse represents a standard DeepSeek API response
// for both chat and reasoning models.
type ChatResponse struct {
	ID      string   `json:"id"`      // unique identifier for the completion
	Object  string   `json:"object"`  // usually "chat.completion"
	Created int64    `json:"created"` // unix timestamp (seconds)
	Model   string   `json:"model"`   // model ID, e.g. "deepseek-chat", "deepseek-r1"
	Choices []Choice `json:"choices"` // one or more generated completions
	Usage   *Usage   `json:"usage"`   // token usage stats (may be nil)
	// Some DeepSeek reasoning models can return extra metadata:
	Reasoning *ReasoningInfo `json:"reasoning,omitempty"`
}

// Choice represents one generated option from the model.
type Choice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`       // assistant message (final answer)
	FinishReason string      `json:"finish_reason"` // "stop", "length", "content_filter", "tool_calls"
}

// Usage provides token accounting.
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// ReasoningInfo is included only for reasoning models (R1 family).
type ReasoningInfo struct {
	// Thinking contains the internal reasoning traces (if requested).
	Thinking string `json:"thinking,omitempty"`
	// BudgetUsed reports how many tokens were consumed for reasoning.
	BudgetUsed int `json:"budget_used,omitempty"`
	// TraceID or similar IDs may be included for debugging.
	TraceID string `json:"trace_id,omitempty"`
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

