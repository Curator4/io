package domain

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/curator4/io/backend/internal/database"
	"github.com/google/uuid"
)

// UserToDB converts a domain User to database User
func UserToDB(u User) database.User {
	return database.User{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// ConversationToDB converts a domain Conversation to database Conversation
func ConversationToDB(c Conversation) database.Conversation {
	return database.Conversation{
		ID:         c.ID,
		Name:       stringToSqlNullString(c.Name),
		CreatedAt:  c.CreatedAt,
		UpdatedAt:  c.UpdatedAt,
		LastUsedAt: ptrToSqlNullTime(c.LastUsedAt),
	}
}

// MessageToDB converts a domain Message to database Message
func MessageToDB(m Message) database.Message {
	// Content needs to be marshaled to JSONB (MessageContent already has json tags)
	contentJSON, _ := json.Marshal(m.Content)

	var userID uuid.NullUUID
	if m.User != nil {
		userID = uuid.NullUUID{UUID: m.User.ID, Valid: true}
	}

	return database.Message{
		ID:             m.ID,
		ConversationID: m.ConversationID,
		UserID:         userID,
		Role:           string(m.Role),
		Content:        contentJSON,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.CreatedAt, // Messages don't get updated, so use CreatedAt
	}
}

// ProviderToDB converts a domain Provider to database Provider
func ProviderToDB(p Provider) database.Provider {
	return database.Provider{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

// ModelToDB converts a domain Model to database Model
func ModelToDB(m Model) database.Model {
	return database.Model{
		ID:          m.ID,
		ProviderID:  m.ProviderID,
		Name:        m.Name,
		Description: stringToSqlNullString(m.Description),
		CreatedAt:   m.CreatedAt,
	}
}

// AIConfigToDB converts a domain AIConfig to database AiConfig
func AIConfigToDB(a AIConfig) database.AiConfig {
	return database.AiConfig{
		ID:           a.ID,
		Name:         a.Name,
		ModelID:      a.Model.ID, // Extract model ID from Model object
		SystemPrompt: stringToSqlNullString(a.SystemPrompt),
		CreatedAt:    a.CreatedAt,
		UpdatedAt:    a.UpdatedAt,
		LastUsedAt:   ptrToSqlNullTime(a.LastUsedAt),
	}
}

// ConversationParticipantToDB converts a domain ConversationParticipant to database
func ConversationParticipantToDB(cp ConversationParticipant) database.ConversationParticipant {
	return database.ConversationParticipant{
		ConversationID: cp.ConversationID,
		UserID:         cp.UserID,
		JoinedAt:       cp.JoinedAt,
	}
}

// Helper functions for nullable types
func ptrToSqlNullTime(t *time.Time) sql.NullTime {
	if t == nil {
		return sql.NullTime{Valid: false}
	}
	return sql.NullTime{Time: *t, Valid: true}
}

func stringToSqlNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}
