package mistral

// Mistral model IDs for use with the Mistral AI API.
// These constants represent model identifiers you can pass to the "model" parameter.
const (
	// ModelMistral7B is the standard open-source Mistral model (~7.3B parameters),
	// released in September 2023. Includes both base and instruct-tuned versions.
	ModelMistral7B = "mistral-7b"

	// ModelMixtral8x7B is a sparse mixture-of-experts model (~46.7B parameters total,
	// ~13B active per token), released April 2024. Offers high performance with efficiency.
	ModelMixtral8x7B = "mixtral-8x7b"

	// ModelMixtral8x7BInstruct is the instruct-tuned version of Mixtral 8x7B.
	ModelMixtral8x7BInstruct = "mixtral-8x7b-instruct"

	// ModelCodestral22B is a code-specialized model (~22B parameters),
	// released May 2024. Licensed for research use.
	ModelCodestral22B = "codestral-22b"

	// ModelCodestralMamba7B is a lighter code model (~7B parameters),
	// released July 2024 under Apache 2.0.
	ModelCodestralMamba7B = "codestral-mamba-7b"

	// ModelMathstral7B is a STEM-focused model (~7B parameters),
	// released July 2024 with extended context (32k tokens) under Apache 2.0.
	ModelMathstral7B = "mathstral-7b"

	// ModelMistralSmall31 is the small multimodal model (~24B parameters),
	// released March 2025 (version 3.1) with enhanced reasoning and image understanding.
	ModelMistralSmall31 = "mistral-small-3.1"

	// ModelMistralMedium3 is the mid-tier model ("Medium 3"), released May 2025,
	// offering Claude-comparable performance at lower cost.
	ModelMistralMedium3 = "mistral-medium-3"

	// ModelMagistralSmall is the open-source reasoning model (chain-of-thought),
	// released June 2025 under Apache 2.0.
	ModelMagistralSmall = "magistral-small"

	// ModelMagistralMedium is the more powerful enterprise reasoning model,
	// also released June 2025 (proprietary/commercial).
	ModelMagistralMedium = "magistral-medium"
)

const (
	// DefaultModel is the default Mistral model used for requests.
	DefaultModel = ModelMistral7B
)

// AllModels is a list of all available Mistral model IDs.
var AllModels = []string{
	ModelMistral7B,
	ModelMixtral8x7B,
	ModelMixtral8x7BInstruct,
	ModelCodestral22B,
	ModelCodestralMamba7B,
	ModelMathstral7B,
	ModelMistralSmall31,
	ModelMistralMedium3,
	ModelMagistralSmall,
	ModelMagistralMedium,
}

// GetAllModels returns a list of all available Mistral model IDs.
func GetAllModels() []string {
	return AllModels
}
