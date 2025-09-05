package chatgpt

// GetAllModels returns a list of all available models.
func GetAllModels() []string {
	return []string{
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
}
