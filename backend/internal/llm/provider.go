package llm

import (
	"context"

	"github.com/curator4/io/backend/internal/domain"
)

// Provider is the interface that all AI providers must implement
type Provider interface {
	// SendMessage sends a list of messages to the AI provider and returns the response content
	SendMessage(ctx context.Context, messages []domain.Message, config domain.AIConfig) (domain.MessageContent, error)
}
