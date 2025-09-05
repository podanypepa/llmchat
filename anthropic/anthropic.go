package anthropic

import "encoding/json"

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

type Message struct {
	Role    Role    `json:"role"` // "user" | "assistant"
	Content Content `json:"content"`
}

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

type Content any

type ContentBlock struct {
	Type   string       `json:"type"`
	Text   string       `json:"text,omitempty"`
	Source *ImageSource `json:"source,omitempty"`
}

type ImageSource struct {
	Type      string  `json:"type"` // "base64" | "url"
	MediaType *string `json:"media_type,omitempty"`
	Data      *string `json:"data,omitempty"`
	URL       *string `json:"url,omitempty"`
}

type Tool struct {
	Name         string            `json:"name"` // povinné
	Description  string            `json:"description,omitempty"`
	InputSchema  JSONSchemaObject  `json:"input_schema"`            // JSON Schema (type: object)
	Type         string            `json:"type,omitempty"`          // "custom" (volitelně)
	CacheControl *ToolCacheControl `json:"cache_control,omitempty"` // volitelné cache breakpoints
}

type ToolCacheControl struct {
	Type string  `json:"type"`          // "ephemeral"
	TTL  *string `json:"ttl,omitempty"` // "5m" | "1h"
}

type JSONSchemaObject struct {
	Type       string                        `json:"type"`
	Properties map[string]JSONSchemaProperty `json:"properties,omitempty"`
	Required   []string                      `json:"required,omitempty"`
}

type JSONSchemaProperty struct {
	Type        string                        `json:"type,omitempty"`
	Description string                        `json:"description,omitempty"`
	Enum        []string                      `json:"enum,omitempty"`
	Items       *JSONSchemaProperty           `json:"items,omitempty"`
	Properties  map[string]JSONSchemaProperty `json:"properties,omitempty"`
}

type ToolChoice struct {
	Type                   ToolChoiceType `json:"type"`                                // "auto" | "any" | "tool" | "none"
	ToolName               string         `json:"name,omitempty"`                      // když Type=="tool"
	DisableParallelToolUse *bool          `json:"disable_parallel_tool_use,omitempty"` // default false
}

type ToolChoiceType string

const (
	ToolChoiceAuto ToolChoiceType = "auto"
	ToolChoiceAny  ToolChoiceType = "any"
	ToolChoiceTool ToolChoiceType = "tool"
	ToolChoiceNone ToolChoiceType = "none"
)

type ThinkingCfg struct {
	Type         string `json:"type"`          // musí být "enabled"
	BudgetTokens int    `json:"budget_tokens"` // >=1024 a < max_tokens
}

type ServiceTier string

const (
	ServiceTierAuto         ServiceTier = "auto"
	ServiceTierStandardOnly ServiceTier = "standard_only"
)

type RequestMeta struct {
	UserID string `json:"user_id,omitempty"`
}

type MCPServer struct {
	Name              string         `json:"name"`
	Type              string         `json:"type"` // "url"
	URL               string         `json:"url"`
	Authorization     *string        `json:"authorization_token,omitempty"`
	ToolConfiguration *MCPToolConfig `json:"tool_configuration,omitempty"`
}

type MCPToolConfig struct {
	Enabled      *bool    `json:"enabled,omitempty"`
	AllowedTools []string `json:"allowed_tools,omitempty"`
}

type Response struct {
	ID         string           `json:"id"`            // např. "msg_013Zva2CMH..."
	Type       string           `json:"type"`          // obvykle "message"
	Role       Role             `json:"role"`          // "assistant"
	Model      string           `json:"model"`         // model, který odpovídal
	Content    []AssistantBlock `json:"content"`       // text a/nebo tool_use bloky
	StopReason string           `json:"stop_reason"`   // "end_turn" | "max_tokens" | "stop_sequence" | "tool_use" ...
	StopSeq    *string          `json:"stop_sequence"` // může být null
	Usage      Usage            `json:"usage"`         // počty tokenů
}

type Usage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

type AssistantBlock struct {
	Type      string          `json:"type"` // "text" | "tool_use"
	Text      string          `json:"text,omitempty"`
	ToolUseID string          `json:"id,omitempty"`
	ToolName  string          `json:"name,omitempty"`
	ToolInput json.RawMessage `json:"input,omitempty"`
}
