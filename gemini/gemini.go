// Package gemini implements a Gemini protocol client.
package gemini

import "errors"

// ChatRequest represents a request to Gemini API.
type ChatRequest struct {
	// One or more content blocks (conversation turns).
	Contents []Content `json:"contents"`

	// Optional generation config (sampling etc.).
	GenerationConfig *GenerationConfig `json:"generationConfig,omitempty"`

	// Optional safety settings.
	SafetySettings []SafetySetting `json:"safetySettings,omitempty"`
}

// Content is a single conversational turn (role + parts).
type Content struct {
	Role  string `json:"role"`  // "user" | "model"
	Parts []Part `json:"parts"` // text, images, etc.
}

// Part is a multimodal input element.
type Part struct {
	// For text input.
	Text string `json:"text,omitempty"`

	// For image input (inline base64 or URI).
	InlineData *InlineData `json:"inlineData,omitempty"`
}

// InlineData is used for binary inputs (like images).
type InlineData struct {
	MimeType string `json:"mimeType"`
	Data     string `json:"data"` // base64 encoded
}

// GenerationConfig controls decoding and output length.
type GenerationConfig struct {
	Temperature     float64  `json:"temperature,omitempty"`
	TopP            float64  `json:"topP,omitempty"`
	TopK            int      `json:"topK,omitempty"`
	MaxOutputTokens int      `json:"maxOutputTokens,omitempty"`
	StopSequences   []string `json:"stopSequences,omitempty"`
}

// SafetySetting controls filtering for specific harm categories.
type SafetySetting struct {
	Category  string `json:"category"`
	Threshold string `json:"threshold"` // e.g. "BLOCK_MEDIUM_AND_ABOVE"
}

// ChatResponse is the top-level response from Gemini API.
type ChatResponse struct {
	Candidates     []Candidate     `json:"candidates"`
	PromptFeedback *PromptFeedback `json:"promptFeedback,omitempty"`
	UsageMetadata  *UsageMetadata  `json:"usageMetadata,omitempty"`
}

// Candidate represents one possible generated answer.
type Candidate struct {
	Content       Content        `json:"content"`
	FinishReason  string         `json:"finishReason,omitempty"` // e.g. "STOP", "MAX_TOKENS", "SAFETY"
	Index         int            `json:"index,omitempty"`
	SafetyRatings []SafetyRating `json:"safetyRatings,omitempty"`
}

// SafetyRating indicates how safe the response was classified.
type SafetyRating struct {
	Category    string `json:"category"`    // e.g. "HARM_CATEGORY_DEROGATORY"
	Probability string `json:"probability"` // e.g. "NEGLIGIBLE", "LOW", "MEDIUM", "HIGH"
}

// PromptFeedback gives feedback if prompt triggered safety blocks.
type PromptFeedback struct {
	SafetyRatings []SafetyRating `json:"safetyRatings,omitempty"`
	BlockReason   string         `json:"blockReason,omitempty"` // e.g. "SAFETY"
}

// UsageMetadata gives token accounting.
type UsageMetadata struct {
	PromptTokenCount     int `json:"promptTokenCount"`
	CandidatesTokenCount int `json:"candidatesTokenCount"`
	TotalTokenCount      int `json:"totalTokenCount"`
}

// ExtractText returns the text of the first candidate's first part.
// It will return an error if no candidate or no text is available.
func (r *ChatResponse) ExtractText() (string, error) {
	if r == nil {
		return "", errors.New("response is nil")
	}
	if len(r.Candidates) == 0 {
		return "", errors.New("no candidates in response")
	}
	c := r.Candidates[0]
	if len(c.Content.Parts) == 0 {
		return "", errors.New("no parts in first candidate")
	}
	// Look for first non-empty text part
	for _, p := range c.Content.Parts {
		if p.Text != "" {
			return p.Text, nil
		}
	}
	return "", errors.New("no text part found in first candidate")
}
