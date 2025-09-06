package perplexity

// Perplexity model IDs (as of September 2025).
const (
	// Sonar base models
	ModelPerplexitySonar    = "sonar"     // default Sonar model
	ModelPerplexitySonarPro = "sonar-pro" // higher-capacity Sonar

	// Reasoning-focused models
	ModelPerplexitySonarReasoning    = "sonar-reasoning"     // optimized for reasoning with search
	ModelPerplexitySonarReasoningPro = "sonar-reasoning-pro" // enhanced reasoning capability

	// Deep research model
	ModelPerplexitySonarDeepResearch = "sonar-deep-research" // expert-level deep research
)

const (
	// DefaultModel is the default Perplexity model used for requests.
	DefaultModel = ModelPerplexitySonar
)

// AllModels is a list of all available Perplexity model IDs.
var AllModels = []string{
	ModelPerplexitySonar,
	ModelPerplexitySonarPro,
	ModelPerplexitySonarReasoning,
	ModelPerplexitySonarReasoningPro,
	ModelPerplexitySonarDeepResearch,
}

// GetAllModels returns a list of all available Perplexity model IDs.
func GetAllModels() []string {
	return AllModels
}
