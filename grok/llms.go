package grok

// Grok model IDs used with the xAI / Grok API.
// These constants represent the official API identifiers for Grok model versions.
const (
	// ModelGrok1 is the original Grok model, released in November 2023 as open-source (Apache-2.0).
	ModelGrok1 = "grok-1"

	// ModelGrok15 is the improved Grok model with extended reasoning and a large context window,
	// released in early 2024.
	ModelGrok15 = "grok-1.5"

	// ModelGrok2 is the Grok model with upgraded reasoning and image generation support,
	// released in August 2024.
	ModelGrok2 = "grok-2"

	// ModelGrok2Mini is a lightweight, faster alternative to grok-2, released in August 2024.
	ModelGrok2Mini = "grok-2-mini"

	// ModelGrok3 is the third-generation Grok model with advanced reasoning ("Think" mode),
	// released in February 2025.
	ModelGrok3 = "grok-3"

	// ModelGrok3Mini is a faster, lightweight variant of Grok 3, released in early 2025.
	ModelGrok3Mini = "grok-3-mini"

	// ModelGrok4 is the flagship Grok model released in July 2025. It supports native tool use,
	// real-time search, and enhanced reasoning capabilities.
	ModelGrok4 = "grok-4"

	// ModelGrok4Heavy is the high-performance “Heavy” variant of Grok 4,
	// offered via SuperGrok subscriptions in mid-2025.
	ModelGrok4Heavy = "grok-4-heavy"
)

const (
	// DefaultModel is the default Grok model used for requests.
	DefaultModel = ModelGrok4
)

// AllModels is a list of all available Grok model IDs.
var AllModels = []string{
	ModelGrok1,
	ModelGrok15,
	ModelGrok2,
	ModelGrok2Mini,
	ModelGrok3,
	ModelGrok3Mini,
	ModelGrok4,
	ModelGrok4Heavy,
}
