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
		ID:         c.ID,
		Name:       sqlNullStringToString(c.Name),
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
		LastUsedAt: sqlNullTimeToPtr(c.LastUsedAt),
	}
}

// MessageFromDB converts a database Message to domain Message
func MessageFromDB(m database.Message) Message {
	// Content is stored as JSONB, extract the string value
	var content string
	// Ignore unmarshal errors - if content is invalid JSON, use empty string
	_ = json.Unmarshal(m.Content, &content)

	return Message{
		ID:             m.ID,
		ConversationID: m.ConversationID,
		UserID:         uuidNullUUIDToPtr(m.UserID),
		Role:           m.Role,
		Content:        content,
		CreatedAt:      m.CreatedAt,
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
		UpdatedAt:   m.UpdatedAt,
	}
}

// AIConfigFromDB converts a database AiConfig to domain AIConfig
func AIConfigFromDB(a database.AiConfig) AIConfig {
	return AIConfig{
		ID:           a.ID,
		Name:         a.Name,
		ModelID:      a.ModelID,
		SystemPrompt: sqlNullStringToString(a.SystemPrompt),
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
		LastUsedAt:   sqlNullTimeToPtr(a.LastUsedAt),
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
func uuidNullUUIDToPtr(u uuid.NullUUID) *uuid.UUID {
	if !u.Valid {
		return nil
	}
	return &u.UUID
}

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
