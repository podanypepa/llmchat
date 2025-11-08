package gemini

import (
	"os"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		apikey string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty api key",
			args: args{
				apikey: "",
			},
			wantErr: true,
		},
		{
			name: "valid api key",
			args: args{
				apikey: os.Getenv("DEEPSEEK_API_KEY"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.apikey, GeminiPro)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("NewClient() got = %v, want non-nil client", got)
			}
		})
	}
}

func TestNewClientWithConfig(t *testing.T) {
	type args struct {
		config *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "empty api key",
			args: args{
				config: &Config{
					APIKey: "",
					Model:  DefaultModel,
				},
			},
			wantErr: true,
		},
		{
			name: "valid api key",
			args: args{
				config: &Config{
					APIKey: os.Getenv("DEEPSEEK_API_KEY"),
					Model:  DefaultModel,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClientWithConfig(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClientWithConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("NewClientWithConfig() got = %v, want non-nil client", got)
			}
		})
	}
}
