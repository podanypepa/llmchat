// Package chatgpt provides a client for interacting with the ChatGPT API.
package chatgpt

const (
	// DefaultapiURL is the default base URL for the ChatGPT API.
	DefaultapiURL = "https://api.openai.com/v1"
)

const (
	// RoleSystem is the system role for chat messages.
	RoleSystem = "system"
	// RoleUser is the user role for chat messages.
	RoleUser = "user"
	// RoleAssistant is the assistant role for chat messages.
	RoleAssistant = "assistant"
)

// ChatMessage represents a message in a chat conversation.
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest represents a request to the ChatGPT API.
type ChatRequest struct {
	Model       string        `json:"model"`
	Messages    []ChatMessage `json:"messages"`
	Temperature float64       `json:"temperature,omitempty"`
	TopP        float64       `json:"top_p,omitempty"`
	MaxTokens   int           `json:"max_tokens,omitempty"`
	User        string        `json:"user,omitempty"`
}

// ChatResponse represents a response from the ChatGPT API.
type ChatResponse struct {
	ID        string               `json:"id"`
	Object    string               `json:"object"`
	CreatedAt int64                `json:"created_at"`
	Choices   []ChatResponseChoice `json:"choices"`
	Usage     ChatResponseUsage    `json:"usage"`
}

// ChatResponseChoice represents a single choice in a ChatResponse.
type ChatResponseChoice struct {
	Index        int         `json:"index"`
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
}

// ChatResponseUsage represents token usage information in a ChatResponse.
type ChatResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
