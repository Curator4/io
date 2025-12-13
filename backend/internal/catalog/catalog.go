package catalog

import (
	"fmt"
	"slices"
)

// ProviderName represents an LLM provider
type ProviderName string

// ModelName represents an LLM model
type ModelName string

// Provider constants
const (
	ProviderOpenAI    ProviderName = "openai"
	ProviderAnthropic ProviderName = "anthropic"
)

// Model constants
const (
	// OpenAI models
	ModelGPT5Nano ModelName = "gpt-5-nano"
	ModelGPT4Nano ModelName = "gpt-4.1-nano"

	// Anthropic models
	ModelClaude35Sonnet ModelName = "claude-3-5-sonnet-20241022"
	ModelClaude35Haiku  ModelName = "claude-3-5-haiku-20241022"
)

// ValidModels maps providers to their supported models
var ValidModels = map[ProviderName][]ModelName{
	ProviderOpenAI:    {ModelGPT5Nano, ModelGPT4Nano},
	ProviderAnthropic: {ModelClaude35Sonnet, ModelClaude35Haiku},
}

// Default configuration values
const (
	DefaultProviderName = ProviderOpenAI
	DefaultModelName    = ModelGPT5Nano
	DefaultConfigName   = "Io"
	DefaultSystemPrompt = `You are Io, a helpful AI assistant.

	Context: You're in a Discord group chat. Messages are formatted as "username: message content" - the username is part of the format, not part of what the user is saying. Multiple people may be talking.

	Guidelines:
	- Match the tone of the conversation - casual when users are casual, professional when they need help.
	- Keep responses concise and Discord-friendly.
	- For technical questions: be precise, show examples, explain clearly.
	- Use natural language - don't force slang or emojis unless it fits the context.
	- IMPORTANT: Do NOT prefix your responses with "io:" or "Io:" - just respond directly.`
)

// GetModelsForProvider returns all valid models for a given provider
func GetModelsForProvider(provider ProviderName) []ModelName {
	return ValidModels[provider]
}

// GetAllProviders returns all supported providers
func GetAllProviders() []ProviderName {
	providers := make([]ProviderName, 0, len(ValidModels))
	for p := range ValidModels {
		providers = append(providers, p)
	}
	return providers
}

// ValidateModel checks if a model is valid for a given provider
func ValidateModel(provider ProviderName, model ModelName) error {
	models, ok := ValidModels[provider]
	if !ok {
		return fmt.Errorf("unsupported provider: %s", provider)
	}

	if !slices.Contains(models, model) {
		return fmt.Errorf("unsupported model %s for provider %s", model, provider)
	}

	return nil
}
