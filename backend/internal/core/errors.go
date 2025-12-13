package core

import "errors"

// Sentinel errors for mapping to gRPC status codes
var (
	// ErrNoConfigsFound indicates no AI configs exist in the database
	ErrNoConfigsFound = errors.New("no ai configs found")

	// ErrProviderNotFound indicates the requested LLM provider is not registered
	ErrProviderNotFound = errors.New("llm provider not found")

	// ErrLLMUnavailable indicates the LLM service failed to respond
	ErrLLMUnavailable = errors.New("llm service unavailable")
)
