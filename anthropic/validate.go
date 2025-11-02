package anthropic

import "fmt"

func validate(r *ChatRequest) error {
	if r.Model == "" {
		return fmt.Errorf("model is required")
	}
	if r.MaxTokens < 1 {
		return fmt.Errorf("max_tokens must be at least 1")
	}
	if len(r.Messages) == 0 {
		return fmt.Errorf("at least one message is required")
	}
	for i, msg := range r.Messages {
		if msg.Role == "" {
			return fmt.Errorf("message %d: role is required", i)
		}
		if msg.Content == nil {
			return fmt.Errorf("message %d: content is required", i)
		}
		switch content := msg.Content.(type) {
		case string:
			if content == "" {
				return fmt.Errorf("message %d: content is required", i)
			}
		case []ContentBlock:
			if len(content) == 0 {
				return fmt.Errorf("message %d: content is required", i)
			}
		default:
			// To support []any and other variations
			if fmt.Sprintf("%v", msg.Content) == "[]" || fmt.Sprintf("%v", msg.Content) == "" {
				return fmt.Errorf("message %d: content is required", i)
			}
		}
	}
	return nil
}
