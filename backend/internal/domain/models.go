package domain

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Conversation represents a chat conversation
type Conversation struct {
	ID         uuid.UUID
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	LastUsedAt *time.Time
}

// Message represents a chat message
type Message struct {
	ID             uuid.UUID
	ConversationID uuid.UUID
	UserID         *uuid.UUID // nil for assistant messages
	Role           string     // "user", "assistant", "system"
	Content        string
	CreatedAt      time.Time
}

// Provider represents an AI provider (OpenAI, Anthropic, etc.)
type Provider struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Model represents an AI model
type Model struct {
	ID          uuid.UUID
	ProviderID  uuid.UUID
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// AIConfig represents an AI configuration
type AIConfig struct {
	ID           uuid.UUID
	Name         string
	ModelID      uuid.UUID
	SystemPrompt string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastUsedAt   *time.Time
}

// ConversationParticipant represents a user's participation in a conversation
type ConversationParticipant struct {
	ConversationID uuid.UUID
	UserID         uuid.UUID
	JoinedAt       time.Time
}
