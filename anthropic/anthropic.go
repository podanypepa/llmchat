package anthropic

// Message represents a message in the conversation.
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Request represents a request to the Anthropic API.
type Request struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens"`
	System      string    `json:"system,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
}

// Content represents the content of a response message.
type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// Response represents a response from the Anthropic API.
type Response struct {
	ID      string    `json:"id"`
	Type    string    `json:"type"`
	Role    string    `json:"role"`
	Model   string    `json:"model"`
	Content []Content `json:"content"`
}
