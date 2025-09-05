package anthropic

const (
	// RoleUser represents a user message role.
	RoleUser = "user"
	// RoleSystem represents a system message role.
	RoleSystem = "system"
)

type Request struct {
	Model       string         `json:"model"`              // např. "claude-3-5-sonnet-20240620"
	Messages    []Message      `json:"messages"`           // konverzační historie
	MaxTokens   int            `json:"max_tokens"`         // max. počet tokenů v odpovědi
	Metadata    map[string]any `json:"metadata,omitempty"` // volitelné metadata
	Stream      bool           `json:"stream,omitempty"`   // zda streamovat odpověď
	StopSeq     []string       `json:"stop_sequences,omitempty"`
	Temperature *float64       `json:"temperature,omitempty"`
	TopP        *float64       `json:"top_p,omitempty"`
	TopK        *int           `json:"top_k,omitempty"`
}

type Message struct {
	Role    string   `json:"role"`    // "user" | "assistant"
	Content []string `json:"content"` // nebo složitější struct pokud chceš podporovat text+obrázky
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
