package deepseek

// DeepSeek model IDs available via the DeepSeek API.
// These constants represent stable model identifiers that can be used with the "model" parameter.
const (
	// ModelDeepSeekChat is the conversational model interface (“chat”) built on DeepSeek-V3.
	ModelDeepSeekChat = "deepseek-chat"

	// ModelDeepSeekV3 is the general-purpose MoE model (V3) optimized for reasoning and long context.
	ModelDeepSeekV3 = "deepseek-v3"

	// ModelDeepSeekV3_0324 is the improved version of V3 released on March 24, 2025
	// with enhanced reasoning and multilingual support.
	ModelDeepSeekV3_0324 = "deepseek-v3-0324"

	// ModelDeepSeekV3_1 is the upgraded V3 model (V3.1) released on August 21, 2025,
	// featuring hybrid “thinking/non-thinking” modes and expanded context up to 128K tokens.
	ModelDeepSeekV3_1 = "deepseek-v3.1"

	// ModelDeepSeekR1 is the reasoning-focused model (R1),
	// trained via reinforcement learning to excel in logic, math, and code.
	ModelDeepSeekR1 = "deepseek-r1"

	// ModelDeepSeekR1_0528 is an improved version of R1 released on May 28, 2025,
	// with deeper inference and higher accuracy benchmarks.
	ModelDeepSeekR1_0528 = "deepseek-r1-0528"

	// ModelDeepSeekR1Zero is the original RL-trained reasoning model without supervised fine-tuning,
	// notable for its emergent reasoning behaviors.
	ModelDeepSeekR1Zero = "deepseek-r1-zero"
)

const (
	// DefaultModel is the default DeepSeek model used for requests.
	DefaultModel = ModelDeepSeekChat
)

// AllModels is a list of all available DeepSeek model IDs.
var AllModels = []string{
	ModelDeepSeekChat,
	ModelDeepSeekV3,
	ModelDeepSeekV3_0324,
	ModelDeepSeekV3_1,
	ModelDeepSeekR1,
	ModelDeepSeekR1_0528,
	ModelDeepSeekR1Zero,
}

// GetAllModels returns a list of all available DeepSeek model IDs.
func GetAllModels() []string {
	return AllModels
}
