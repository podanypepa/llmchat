package chatgpt

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		req     *ChatRequest
		wantErr bool
	}{
		{
			name:    "valid request",
			req:     &ChatRequest{Model: Gpt4OMini, Messages: []ChatMessage{{Role: "user", Content: "Hello"}}},
			wantErr: false,
		},
		{
			name:    "missing model",
			req:     &ChatRequest{Messages: []ChatMessage{{Role: "user", Content: "Hello"}}},
			wantErr: true,
		},
		{
			name:    "no messages",
			req:     &ChatRequest{Model: Gpt4OMini},
			wantErr: true,
		},
		{
			name:    "message missing role",
			req:     &ChatRequest{Model: Gpt4OMini, Messages: []ChatMessage{{Content: "Hello"}}},
			wantErr: true,
		},
		{
			name:    "message missing content",
			req:     &ChatRequest{Model: Gpt4OMini, Messages: []ChatMessage{{Role: "user"}}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate(tt.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
