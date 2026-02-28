package domain

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleSystem    Role = "system"
	RoleDeveloper Role = "developer"
)

// MediaItem represents a single media attachment (image, video, file, etc.)
type MediaItem struct {
	Type     string `json:"type"`                // "image", "video", "file", "audio"
	URL      string `json:"url"`                 // URL to the media
	FileName string `json:"file_name,omitempty"` // Original filename if applicable
}

// MessageContent represents the content of a message, supporting text and multiple media attachments
type MessageContent struct {
	Text  string      `json:"text,omitempty"`
	Media []MediaItem `json:"media,omitempty"`
}

// User represents a user in the system
type User struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Conversation represents a chat conversation
type Conversation struct {
	ID           uuid.UUID
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	LastUsedAt   *time.Time
	Participants []ConversationParticipant
}

// Message represents a chat message
type Message struct {
	ID             uuid.UUID
	ConversationID uuid.UUID
	User           *User // Full user object (nil for assistant messages)
	Role           Role  // "user", "assistant", "system"
	Content        MessageContent
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
	Provider    Provider // Full provider object
	Name        string
	Description string
	CreatedAt   time.Time
}

// AIConfig represents an AI configuration
type AIConfig struct {
	ID           uuid.UUID
	Name         string
	Model        Model // Full model object instead of just ID
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

// ConversationLifecycle tracks conversation state changes so can send to front end
type ConversationLifecycle struct {
	IsNewConversation bool
	ConversationID    uuid.UUID
	ConversationName  string
	StartedAt         time.Time
}

// DiscordAction defines ai actions that can be triggered in discord by ai
type DiscordAction struct {
	Type            ActionType
	Reaction        *Reaction
	WebSearchResult *WebSearchResult
	ImageGeneration *ImageGeneration
	CodeInterpreter *CodeInterpreterResult
}

type ActionType string

const (
	ActionTypeReaction        ActionType = "reaction"
	ActionTypeWebSearch       ActionType = "web_search"
	ActionTypeImageGeneration ActionType = "image_generation"
	ActionTypeCodeInterpreter ActionType = "code_interpreter"
)

type Reaction struct {
	Emoji string
}

type WebSearchResult struct {
	Queries []string
	Results []WebSearchItem
}

type WebSearchItem struct {
	Title   string
	URL     string
	Snippet string
}

type ImageGeneration struct {
	ImageURL string
}

type CodeInterpreterResult struct {
	Code    string
	Outputs []string
}

type LLMResponse struct {
	Content MessageContent
	Actions []DiscordAction
}
