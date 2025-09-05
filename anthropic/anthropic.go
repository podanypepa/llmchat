package anthropic

import "encoding/json"

// Request represents a request to the Messages API.
type Request struct {
	Model       string       `json:"model"`
	Messages    []Message    `json:"messages"`
	MaxTokens   int          `json:"max_tokens,omitempty"`
	System      any          `json:"system,omitempty"`
	Metadata    *RequestMeta `json:"metadata,omitempty"`
	StopSeq     []string     `json:"stop_sequences,omitempty"`
	Temperature *float64     `json:"temperature,omitempty"`
	TopP        *float64     `json:"top_p,omitempty"`
	TopK        *int         `json:"top_k,omitempty"`
	Stream      bool         `json:"stream,omitempty"`
	ServiceTier *ServiceTier `json:"service_tier,omitempty"`
	Tools       []Tool       `json:"tools,omitempty"`
	ToolChoice  *ToolChoice  `json:"tool_choice,omitempty"`
	Thinking    *ThinkingCfg `json:"thinking,omitempty"`
	MCPServers  []MCPServer  `json:"mcp_servers,omitempty"`
	Container   *string      `json:"container,omitempty"`
}

// Message represents a single message in the conversation.
type Message struct {
	Role    Role    `json:"role"` // "user" | "assistant"
	Content Content `json:"content"`
}

// Role represents the role of the message sender.
type Role string

const (
	// RoleUser is a message from the user.
	RoleUser Role = "user"
	// RoleAssistant is a message from the assistant (model).
	RoleAssistant Role = "assistant"
)

// Content can be a string (text) or an array of ContentBlock (text and/or image).
type Content any

// ContentBlock represents a block of content, which can be text or an image.
type ContentBlock struct {
	Type   string       `json:"type"`
	Text   string       `json:"text,omitempty"`
	Source *ImageSource `json:"source,omitempty"`
}

// ImageSource represents an image provided as a base64 string or a URL.
type ImageSource struct {
	Type      string  `json:"type"` // "base64" | "url"
	MediaType *string `json:"media_type,omitempty"`
	Data      *string `json:"data,omitempty"`
	URL       *string `json:"url,omitempty"`
}

// Tool represents a tool that the model can use.
type Tool struct {
	Name         string            `json:"name"` // povinné
	Description  string            `json:"description,omitempty"`
	InputSchema  JSONSchemaObject  `json:"input_schema"`
	Type         string            `json:"type,omitempty"`
	CacheControl *ToolCacheControl `json:"cache_control,omitempty"`
}

// ToolCacheControl defines caching behavior for tool results.
type ToolCacheControl struct {
	Type string  `json:"type"`          // "ephemeral"
	TTL  *string `json:"ttl,omitempty"` // "5m" | "1h"
}

// JSONSchemaObject defines a JSON schema object.
type JSONSchemaObject struct {
	Type       string                        `json:"type"`
	Properties map[string]JSONSchemaProperty `json:"properties,omitempty"`
	Required   []string                      `json:"required,omitempty"`
}

// JSONSchemaProperty defines a property in a JSON schema.
type JSONSchemaProperty struct {
	Type        string                        `json:"type,omitempty"`
	Description string                        `json:"description,omitempty"`
	Enum        []string                      `json:"enum,omitempty"`
	Items       *JSONSchemaProperty           `json:"items,omitempty"`
	Properties  map[string]JSONSchemaProperty `json:"properties,omitempty"`
}

// ToolChoice specifies how the model can use tools.
type ToolChoice struct {
	Type                   ToolChoiceType `json:"type"`
	ToolName               string         `json:"name,omitempty"`
	DisableParallelToolUse *bool          `json:"disable_parallel_tool_use,omitempty"`
}

// ToolChoiceType defines how the model can use tools.
type ToolChoiceType string

const (
	// ToolChoiceAuto lets the model decide when to use tools.
	ToolChoiceAuto ToolChoiceType = "auto"
	// ToolChoiceAny allows the model to use any of the provided tools.
	ToolChoiceAny ToolChoiceType = "any"
	// ToolChoiceTool forces the model to use a specific tool.
	ToolChoiceTool ToolChoiceType = "tool"
	// ToolChoiceNone prevents the model from using any tools.
	ToolChoiceNone ToolChoiceType = "none"
)

// ThinkingCfg enables the model to use "thinking" mode with a specified token budget.
type ThinkingCfg struct {
	Type         string `json:"type"`
	BudgetTokens int    `json:"budget_tokens"` // >=1024 a < max_tokens
}

// ServiceTier specifies the service tier for the request.
type ServiceTier string

const (
	// ServiceTierAuto lets the API choose the best tier based on the model.
	ServiceTierAuto ServiceTier = "auto"
	// ServiceTierStandardOnly restricts the request to standard tier models only.
	ServiceTierStandardOnly ServiceTier = "standard_only"
)

// RequestMeta contains metadata about the request, such as user ID.
type RequestMeta struct {
	UserID string `json:"user_id,omitempty"`
}

// MCPServer represents a custom Model Control Plane server configuration.
type MCPServer struct {
	Name              string         `json:"name"`
	Type              string         `json:"type"` // "url"
	URL               string         `json:"url"`
	Authorization     *string        `json:"authorization_token,omitempty"`
	ToolConfiguration *MCPToolConfig `json:"tool_configuration,omitempty"`
}

// MCPToolConfig specifies which tools are allowed to be used by the model when connected to this MCP server.
type MCPToolConfig struct {
	Enabled      *bool    `json:"enabled,omitempty"`
	AllowedTools []string `json:"allowed_tools,omitempty"`
}

// Response represents the response from the Messages API.
type Response struct {
	ID         string           `json:"id"`
	Type       string           `json:"type"`
	Role       Role             `json:"role"`
	Model      string           `json:"model"`
	Content    []AssistantBlock `json:"content"`
	StopReason string           `json:"stop_reason"`
	StopSeq    *string          `json:"stop_sequence"`
	Usage      Usage            `json:"usage"`
}

// Usage contains token usage statistics.
type Usage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

// AssistantBlock represents a block of content in the assistant's response.
type AssistantBlock struct {
	Type      string          `json:"type"` // "text" | "tool_use"
	Text      string          `json:"text,omitempty"`
	ToolUseID string          `json:"id,omitempty"`
	ToolName  string          `json:"name,omitempty"`
	ToolInput json.RawMessage `json:"input,omitempty"`
}
