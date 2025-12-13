package domain

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/curator4/io/backend/internal/database"
	"github.com/google/uuid"
)

// UserFromDB converts a database User to domain User
func UserFromDB(u database.User) User {
	return User{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// ConversationFromDB converts a database Conversation to domain Conversation
func ConversationFromDB(c database.Conversation) Conversation {
	return Conversation{
		ID:           c.ID,
		Name:         sqlNullStringToString(c.Name),
		CreatedAt:    c.CreatedAt,
		UpdatedAt:    c.UpdatedAt,
		LastUsedAt:   sqlNullTimeToPtr(c.LastUsedAt),
		Participants: []ConversationParticipant{}, // Initialize empty, load separately if needed
	}
}

// MessageFromDB converts a database query result to domain Message
func MessageFromDB(row database.GetMessagesByConversationRow) Message {
	// Content is stored as JSONB
	var content MessageContent
	// Ignore unmarshal errors - if content is invalid JSON, use empty content
	_ = json.Unmarshal(row.Content, &content)

	var user *User
	if row.User.ID != uuid.Nil {
		u := UserFromDB(row.User)
		user = &u
	}

	return Message{
		ID:             row.ID,
		ConversationID: row.ConversationID,
		User:           user,
		Role:           Role(row.Role),
		Content:        content,
		CreatedAt:      row.CreatedAt,
	}
}

// ProviderFromDB converts a database Provider to domain Provider
func ProviderFromDB(p database.Provider) Provider {
	return Provider{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ModelFromDB converts a database Model to domain Model
func ModelFromDB(m database.Model) Model {
	return Model{
		ID:          m.ID,
		ProviderID:  m.ProviderID,
		Name:        m.Name,
		Description: sqlNullStringToString(m.Description),
		CreatedAt:   m.CreatedAt,
	}
}

// AIConfigFromDB converts a database query result to domain AIConfig
func AIConfigFromDB(row database.GetAIConfigByIDRow) AIConfig {
	// Construct model with embedded provider
	model := Model{
		ID:          row.Model.ID,
		Provider:    ProviderFromDB(row.Provider),
		Name:        row.Model.Name,
		Description: sqlNullStringToString(row.Model.Description),
		CreatedAt:   row.Model.CreatedAt,
	}

	return AIConfig{
		ID:           row.ID,
		Name:         row.Name,
		Model:        model,
		SystemPrompt: sqlNullStringToString(row.SystemPrompt),
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
		LastUsedAt:   sqlNullTimeToPtr(row.LastUsedAt),
	}
}

// ConversationParticipantFromDB converts a database ConversationParticipant to domain
func ConversationParticipantFromDB(cp database.ConversationParticipant) ConversationParticipant {
	return ConversationParticipant{
		ConversationID: cp.ConversationID,
		UserID:         cp.UserID,
		JoinedAt:       cp.JoinedAt,
	}
}

// Helper functions for nullable types
func sqlNullTimeToPtr(t sql.NullTime) *time.Time {
	if !t.Valid {
		return nil
	}
	return &t.Time
}

func sqlNullStringToString(s sql.NullString) string {
	if !s.Valid {
		return ""
	}
	return s.String
}
