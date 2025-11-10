package llm

// ChatMessage represents a generic message in a chat conversation.
type ChatMessage struct {
	Role    string
	Content string
}

// Request represents a generic request to an LLM API.
type Request struct {
	Model    string
	Messages []ChatMessage
}

// Usage represents token usage information.
type Usage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

// Metadata contains metadata about the response.
type Metadata struct {
	Usage Usage
}

// Response represents a generic response from an LLM API.
type Response struct {
	Content  string
	Metadata Metadata
}

// StreamChunk is a single chunk of a streamed completion.
type StreamChunk struct {
	Content string
	Err     error
}
