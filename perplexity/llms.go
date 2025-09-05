package perplexity

// Perplexity (Sonar) model IDs for use with the Perplexity API.
// These constants represent the internal Sonar model identifiers you can pass to the "model" parameter.
const (
	// ModelPerplexitySonar is the default model used when none is specified.
	ModelPerplexitySonar = "perplexity/sonar"

	// ModelPerplexitySonarPro is a higher-capacity Sonar model optimized for real-time reasoning.
	ModelPerplexitySonarPro = "perplexity/sonar-pro"

	// ModelPerplexitySonarReasoning is a fast Sonar model tailored for quick reasoning with search.
	ModelPerplexitySonarReasoning = "perplexity/sonar-reasoning"

	// ModelPerplexitySonarReasoningPro is an enhanced reasoning model with extended capability.
	ModelPerplexitySonarReasoningPro = "perplexity/sonar-reasoning-pro"

	// ModelPerplexitySonarDeepResearch is the deep-research Sonar model for expert-level, exhaustive analysis.
	ModelPerplexitySonarDeepResearch = "perplexity/sonar-deep-research"
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
