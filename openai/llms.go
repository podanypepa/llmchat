package chatgpt

const (
	// Gpt5 is the latest and most capable model.
	Gpt5 = "gpt-5"
	// Gpt5mini is a smaller and faster version of Gpt5.
	Gpt5mini = "gpt-5-mini"
	// Gpt5nano is the smallest and fastest version of Gpt5.
	Gpt5nano = "gpt-5-nano"
	// Gpt4_1 is an advanced version of Gpt4 with improved capabilities.
	Gpt4_1 = "gpt-4.1"
	// Gpt4_1Mini is a smaller and faster version of Gpt4_1.
	Gpt4_1Mini = "gpt-4.1-mini"
	// Gpt4_1Nano is the smallest and fastest version of Gpt4_1.
	Gpt4_1Nano = "gpt-4.1-nano"
	// GptO3DeepResearch is a model optimized for deep research tasks.
	GptO3DeepResearch = "o3-deep-research"
	// GptO4MiniDeepResearch is a smaller version of GptO3DeepResearch.
	GptO4MiniDeepResearch = "o4-mini-deep-research"
	// GptO3Pro is a professional-grade model for advanced applications.
	GptO3Pro = "o3-pro"
	// GptO3 is a general-purpose model.
	GptO3 = "o3"
	// GptO4Mini is a smaller and faster version of GptO4.
	GptO4Mini = "o4-mini"
	// Gpt4O is a variant of Gpt4 optimized for various tasks.
	Gpt4O = "gpt-4o"
	// Gpt4OMini is a smaller and faster version of Gpt4O.
	Gpt4OMini = "gpt-4o-mini"
)

// DefaultModel is the default model used by the client.
var DefaultModel = Gpt4OMini

// AllModels is a list of all available models.
var AllModels = []string{
	Gpt5,
	Gpt5mini,
	Gpt5nano,
	Gpt4_1,
	Gpt4_1Mini,
	Gpt4_1Nano,
	GptO3DeepResearch,
	GptO4MiniDeepResearch,
	GptO3Pro,
	GptO3,
	GptO4Mini,
	Gpt4O,
	Gpt4OMini,
}
