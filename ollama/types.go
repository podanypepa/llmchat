// Package ollama provides a client for interacting with the Ollama API.
package ollama

const (
	// DefaultBaseURL is the default base URL for the Ollama API.
	DefaultBaseURL = "http://localhost:11434"
)

// ChatMessage represents a message in a chat conversation.
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest represents a request to the Ollama API.
type ChatRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
	Stream   bool          `json:"stream"`
}

// ChatResponse represents a response from the Ollama API.
type ChatResponse struct {
	Model           string      `json:"model"`
	CreatedAt       string      `json:"created_at"`
	Message         ChatMessage `json:"message"`
	Done            bool        `json:"done"`
	TotalDuration   int64       `json:"total_duration"`
	LoadDuration    int64       `json:"load_duration"`
	PromptEvalCount int         `json:"prompt_eval_count"`
	EvalCount       int         `json:"eval_count"`
	EvalDuration    int64       `json:"eval_duration"`
}

// StreamCompletionChunk is a single chunk of a streamed completion.
type StreamCompletionChunk ChatResponse
