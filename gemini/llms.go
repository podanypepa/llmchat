package gemini

// Gemini model IDs for use with the Google Gemini API.
// These constants represent official model identifiers you can use in the `model` parameter.
const (
	// GeminI2_5Pro is Gemini 2.5 Pro — the most capable variant with enhanced reasoning and "Deep Think" mode.
	GeminI2_5Pro = "gemini-2.5-pro"

	// GeminI2_5Flash is Gemini 2.5 Flash — optimized for low-latency and cost-optimized multimodal usage.
	GeminI2_5Flash = "gemini-2.5-flash"

	// GeminI2_5FlashLite is Gemini 2.5 Flash-Lite — cost-efficient model for high throughput, introduced mid-2025.
	GeminI2_5FlashLite = "gemini-2.5-flash-lite"

	// GeminI2_0Flash is Gemini 2.0 Flash — fast and multimodal, good for general-purpose use cases.
	GeminI2_0Flash = "gemini-2.0-flash"

	// GeminI2_0FlashLite is Gemini 2.0 Flash-Lite — highly cost-effective, minimal-latency variant.
	GeminI2_0FlashLite = "gemini-2.0-flash-lite"

	// GeminI1_5Pro is Gemini 1.5 Pro — supports extremely large context (up to 1 million token window), now legacy.
	GeminI1_5Pro = "gemini-1.5-pro"

	// GeminI1_5Flash is Gemini 1.5 Flash — legacy flash variant for general tasks.
	GeminI1_5Flash = "gemini-1.5-flash"

	// GeminI1_0Ultra is Gemini 1.0 Ultra — highly capable original variant, discontinued.
	GeminI1_0Ultra = "gemini-1.0-ultra"

	// GeminI1_0Pro is Gemini 1.0 Pro — general-purpose original variant, discontinued.
	GeminI1_0Pro = "gemini-1.0-pro"

	// GeminI1_0Nano is Gemini 1.0 Nano — designed for on-device inference, discontinued.
	GeminI1_0Nano = "gemini-1.0-nano"
)

const (
	// DefaultModel is the default Gemini model used for requests.
	DefaultModel = GeminI2_5Pro
)

// AllModels is a list of all available Gemini model IDs.
var AllModels = []string{
	GeminI2_5Pro,
	GeminI2_5Flash,
	GeminI2_5FlashLite,
	GeminI2_0Flash,
	GeminI2_0FlashLite,
	GeminI1_5Pro,
	GeminI1_5Flash,
	GeminI1_0Ultra,
	GeminI1_0Pro,
	GeminI1_0Nano,
}

// GetAllModels returns a list of all available Gemini model IDs.
func GetAllModels() []string {
	return AllModels
}
