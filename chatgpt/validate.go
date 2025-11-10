package chatgpt

import "fmt"

func validate(r *ChatRequest) error {
	if r.Model == "" {
		return fmt.Errorf("model is required")
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
		// Also check for empty string if content is of type string
		if contentStr, ok := msg.Content.(string); ok && contentStr == "" {
			return fmt.Errorf("message %d: content cannot be an empty string", i)
		}
	}
	return nil
}
