package chatgpt

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestResponses(t *testing.T) {
	_ = godotenv.Load()
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		t.Skip("OPENAI_API_KEY not set, skipping test")
	}

	client, err := NewClient(apiKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	req := &ResponsesRequest{
		Model: "gpt-4.1-mini",
		Instructions: "Jsi český asistent. " +
			"Pokud potřebuješ aktuální datum a čas, použij MCP server označený 'switchly-mcp'.",
		Input: "Jaké je právě teď aktuální datum a čas? Pokud je to možné, zjisti to přes MCP.",
		Tools: []MCPTool{
			{
				Type:              "mcp",
				ServerLabel:       "switchly-mcp",
				ServerURL:         "https://mcp.switchly.ai/mcp",
				ServerDescription: "Switchly MCP server s různými nástroji.",
				AllowedTools: []string{
					"now",
					"echo",
				},
				RequireApproval: "never",
				Authorization:   os.Getenv("MCP_API_KEY"),
			},
		},
	}

	resp, err := client.Responses(context.Background(), req)
	if err != nil {
		t.Fatalf("Responses call failed: %v", err)
	}

	if resp == nil {
		t.Fatal("Received nil response")
	}

	fmt.Println("out:", ExtractResponsesAssistantText(resp))
	fmt.Println("usage:", resp.Usage.InputTokens, resp.Usage.OutputTokens, resp.Usage.TotalTokens)
}
