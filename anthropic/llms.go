package anthropic

// Anthropic model IDs used with the Messages API.
// These constants represent stable API identifiers for each Claude model version.
const (
	// ModelClaudeSonnet4 is the balanced fourth-generation Claude model (Sonnet),
	// released in May 2025. It offers a good trade-off between performance and efficiency.
	ModelClaudeSonnet4 = "claude-sonnet-4-20250514"

	// ModelClaudeOpus41 is the most capable Claude model of the 4.1 generation (Opus),
	// released in August 2025. Best for complex and demanding tasks.
	ModelClaudeOpus41 = "claude-opus-4-1-20250805"

	// ModelClaudeOpus4 is the fourth-generation Claude model (Opus),
	// released in May 2025. Predecessor of the 4.1 Opus variant.
	ModelClaudeOpus4 = "claude-opus-4"

	// ModelClaude35Sonnet is the improved 3.5-generation Claude model (Sonnet),
	// released in October 2024. A strong general-purpose option.
	ModelClaude35Sonnet = "claude-3-5-sonnet-20241022"

	// ModelClaude35Haiku is the fast and cost-efficient 3.5-generation Claude model (Haiku),
	// released in October 2024. Best for lightweight and quick tasks.
	ModelClaude35Haiku = "claude-3-5-haiku-20241022"

	// ModelClaude3Sonnet is the third-generation Claude model (Sonnet),
	// released in March 2024. Balanced between quality and performance.
	ModelClaude3Sonnet = "claude-3-sonnet-20240307"

	// ModelClaude3Haiku is the fast and lightweight third-generation Claude model (Haiku),
	// released in March 2024. Designed for simple queries and speed.
	ModelClaude3Haiku = "claude-3-haiku-20240307"

	// ModelClaude3Opus is the most powerful third-generation Claude model (Opus),
	// released in February 2024. Suitable for more complex reasoning tasks.
	ModelClaude3Opus = "claude-3-opus-20240229"
)

const (
	// DefaultModel is the default model used by the client.
	DefaultModel = ModelClaudeSonnet4
)

// AllModels is a list of all available Anthropic Claude models.
var AllModels = []string{
	ModelClaudeSonnet4,
	ModelClaudeOpus41,
	ModelClaudeOpus4,
	ModelClaude35Sonnet,
	ModelClaude35Haiku,
	ModelClaude3Sonnet,
	ModelClaude3Haiku,
	ModelClaude3Opus,
}

// GetAllModels returns a list of all available Anthropic Claude models.
func GetAllModels() []string {
	return AllModels
}
