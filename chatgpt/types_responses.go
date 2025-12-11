package chatgpt

import "encoding/json"

// ResponseChunk represents a chunk of the response content.
type ResponseChunk struct {
	Type string          `json:"type"`
	Text json.RawMessage `json:"text,omitempty"`
}

// ResponseOutput represents the output of a response with role and content.
type ResponseOutput struct {
	Role    string          `json:"role"`
	Content []ResponseChunk `json:"content"`
}

// MCPTool represents a tool configuration for MCP (Model Control Protocol).
type MCPTool struct {
	Type              string   `json:"type"`                 // "mcp"
	ServerLabel       string   `json:"server_label"`         // remote MCP server label
	ServerURL         string   `json:"server_url,omitempty"` // remote MCP server URL
	ServerDescription string   `json:"server_description,omitempty"`
	Authorization     string   `json:"authorization,omitempty"`    // mcp token
	AllowedTools      []string `json:"allowed_tools,omitempty"`    // list of allowed tool names
	RequireApproval   string   `json:"require_approval,omitempty"` // "never"
}

// ResponsesRequest represents a request to the responses endpoint.
type ResponsesRequest struct {
	Model        string    `json:"model"`
	Instructions string    `json:"instructions,omitempty"`
	Input        any       `json:"input"`
	Tools        []MCPTool `json:"tools,omitempty"`
}

// ResponseTextContent represents the text content in a response.
type ResponseTextContent struct {
	Value string `json:"value"`
}

// ResponsesResponse represents the response from the responses endpoint.
type ResponsesResponse struct {
	ID                 string         `json:"id"`
	Object             string         `json:"object"`
	CreatedAt          int64          `json:"created_at"`
	Status             string         `json:"status"`
	Background         bool           `json:"background"`
	Billing            BillingInfo    `json:"billing"`
	Error              *ResponseError `json:"error"`
	IncompleteDetails  any            `json:"incomplete_details"`
	Instructions       string         `json:"instructions"`
	MaxOutputTokens    *int           `json:"max_output_tokens"`
	MaxToolCalls       *int           `json:"max_tool_calls"`
	Model              string         `json:"model"`
	Output             []OutputItem   `json:"output"`
	ParallelToolCalls  bool           `json:"parallel_tool_calls"`
	PreviousResponseID *string        `json:"previous_response_id"`
	PromptCacheKey     *string        `json:"prompt_cache_key"`
	PromptCacheRet     *int           `json:"prompt_cache_retention"`
	Reasoning          ReasoningInfo  `json:"reasoning"`
	SafetyIdentifier   any            `json:"safety_identifier"`
	ServiceTier        string         `json:"service_tier"`
	Store              bool           `json:"store"`
	Temperature        float64        `json:"temperature"`
	Text               TextSettings   `json:"text"`
	ToolChoice         any            `json:"tool_choice"` // string | object
	Tools              []ToolConfig   `json:"tools"`
	TopLogprobs        int            `json:"top_logprobs"`
	TopP               float64        `json:"top_p"`
	Truncation         string         `json:"truncation"`
	Usage              UsageInfo      `json:"usage"`
	User               any            `json:"user"`
	Metadata           map[string]any `json:"metadata"`
}

// BillingInfo represents billing information in the response.
type BillingInfo struct {
	Payer string `json:"payer"`
}

// ResponseError represents an error in the response.
type ResponseError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    string `json:"code"`
}

// ReasoningInfo represents reasoning information in the response.
type ReasoningInfo struct {
	Effort  any `json:"effort"`
	Summary any `json:"summary"`
}

// TextSettings represents text settings in the response.
type TextSettings struct {
	Format struct {
		Type string `json:"type"`
	} `json:"format"`
	Verbosity string `json:"verbosity"`
}

// UsageInfo represents usage information in the response.
type UsageInfo struct {
	InputTokens        int `json:"input_tokens"`
	InputTokensDetails struct {
		CachedTokens int `json:"cached_tokens"`
	} `json:"input_tokens_details"`
	OutputTokens        int `json:"output_tokens"`
	OutputTokensDetails struct {
		ReasoningTokens int `json:"reasoning_tokens"`
	} `json:"output_tokens_details"`
	TotalTokens int `json:"total_tokens"`
}

// ToolConfig represents the configuration of a tool in the response.
type ToolConfig struct {
	Type            string   `json:"type"` // "mcp"
	AllowedTools    []string `json:"allowed_tools"`
	Authorization   string   `json:"authorization"`
	Headers         any      `json:"headers"`
	RequireApproval string   `json:"require_approval"`
	ServerDesc      string   `json:"server_description"`
	ServerLabel     string   `json:"server_label"`
	ServerURL       string   `json:"server_url"`
}

// OutputItem represents an item in the output array of the response.
type OutputItem struct {
	ID                string               `json:"id"`
	Type              string               `json:"type"` // "mcp_list_tools" | "mcp_call" | "message"
	Status            string               `json:"status,omitempty"`
	ServerLabel       string               `json:"server_label,omitempty"`
	Tools             []MCPToolDescription `json:"tools,omitempty"`
	ApprovalRequestID *string              `json:"approval_request_id,omitempty"`
	Arguments         string               `json:"arguments,omitempty"` // JSON string
	CallError         *ResponseError       `json:"error,omitempty"`
	Name              string               `json:"name,omitempty"`
	MCPOutput         string               `json:"output,omitempty"`

	// message
	Role    string           `json:"role,omitempty"`
	Content []MessageContent `json:"content,omitempty"`
}

// MCPToolDescription represents the description of a tool in MCP.
type MCPToolDescription struct {
	Annotations struct {
		ReadOnly bool `json:"read_only"`
	} `json:"annotations"`
	Description string     `json:"description"`
	InputSchema JSONSchema `json:"input_schema"`
	Name        string     `json:"name"`
}

// JSONSchema represents a JSON schema for tool input.
type JSONSchema struct {
	Type                 string                        `json:"type"`
	Required             []string                      `json:"required"`
	Properties           map[string]JSONSchemaProperty `json:"properties"`
	AdditionalProperties bool                          `json:"additionalProperties"`
}

// JSONSchemaProperty represents a property in a JSON schema.
type JSONSchemaProperty struct {
	Type        string `json:"type"`
	Description string `json:"description"`
}

// MessageContent represents the content of a message in the response.
type MessageContent struct {
	Type        string          `json:"type"`
	Annotations []any           `json:"annotations"`
	Logprobs    []any           `json:"logprobs"`
	Text        json.RawMessage `json:"text"`
}

// ExtractResponsesAssistantText extracts the assistant's text response from a ResponsesResponse.
func ExtractResponsesAssistantText(resp *ResponsesResponse) string {
	for _, item := range resp.Output {
		if item.Type != "message" || item.Role != "assistant" {
			continue
		}

		for _, c := range item.Content {
			if c.Type != "output_text" || len(c.Text) == 0 {
				continue
			}

			var obj ResponseTextContent
			if err := json.Unmarshal(c.Text, &obj); err == nil && obj.Value != "" {
				return obj.Value
			}

			var s string
			if err := json.Unmarshal(c.Text, &s); err == nil && s != "" {
				return s
			}
		}
	}
	return ""
}
