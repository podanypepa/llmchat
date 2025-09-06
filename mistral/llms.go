package mistral

// Mistral model IDs for use with the Mistral AI API.
// These constants represent model identifiers you can pass to the "model" parameter.
const (
	// Open-source models
	ModelMistralSmallLatest = "mistral-small-latest" // small instruct model, 2024–2025
	ModelMistralSmall2312   = "mistral-small-2312"   // alias: Open-Mixtral-8x7B
	ModelMistralTiny2312    = "mistral-tiny-2312"    // alias: Open-Mistral-7B

	// Premier models
	ModelMistralMedium3     = "mistral-medium-3"     // released May 2025
	ModelMistralMedium31    = "mistral-medium-3.1"   // extended Medium variant
	ModelMistralLargeLatest = "mistral-large-latest" // flagship dense model (2402 release)

	// Reasoning models
	ModelMagistralSmall  = "magistral-small"  // released June 2025
	ModelMagistralMedium = "magistral-medium" // reasoning-focused, more powerful

	// Coding model
	ModelCodestral2501 = "codestral-25.01" // code-specialized model (Jan 2025)

	// High-performance model (Vertex AI release)
	ModelMistralLarge2411 = "mistral-large-24.11"
)

const (
	// DefaultModel is the default Mistral model used for requests.
	DefaultModel = ModelMistralSmallLatest
)

// AllModels is a list of all available Mistral model IDs.
var AllModels = []string{
	ModelMistralSmallLatest,
	ModelMistralSmall2312,
	ModelMistralTiny2312,
	ModelMistralMedium3,
	ModelMistralMedium31,
	ModelMistralLargeLatest,
	ModelMagistralSmall,
	ModelMagistralMedium,
	ModelCodestral2501,
	ModelMistralLarge2411,
}

// GetAllModels returns a list of all available Mistral model IDs.
func GetAllModels() []string {
	return AllModels
}
