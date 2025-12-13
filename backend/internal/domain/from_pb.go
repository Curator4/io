package domain

import (
	pb "github.com/curator4/io/backend/internal/proto"
	"github.com/google/uuid"
)

// MessageContentFromPb converts protobuf MessageContent to domain MessageContent
func MessageContentFromPb(c *pb.MessageContent) MessageContent {
	if c == nil {
		return MessageContent{}
	}

	content := MessageContent{
		Text: c.Text,
	}

	// Convert media items
	if len(c.Media) > 0 {
		content.Media = make([]MediaItem, len(c.Media))
		for i, item := range c.Media {
			content.Media[i] = MediaItem{
				Type:     item.Type,
				URL:      item.Url,
				FileName: item.FileName,
			}
		}
	}

	return content
}

// UserFromPb converts a protobuf User to domain User
func UserFromPb(u *pb.User) User {
	return User{
		ID:        uuid.MustParse(u.Id),
		Name:      u.Name,
		CreatedAt: u.CreatedAt.AsTime(),
		UpdatedAt: u.UpdatedAt.AsTime(),
	}
}

// ConversationFromPb converts a protobuf Conversation to domain Conversation
func ConversationFromPb(c *pb.Conversation) Conversation {
	conv := Conversation{
		ID:        uuid.MustParse(c.Id),
		Name:      c.Name,
		CreatedAt: c.CreatedAt.AsTime(),
		UpdatedAt: c.UpdatedAt.AsTime(),
	}
	return conv
}

// MessageFromPb converts a protobuf Message to domain Message
func MessageFromPb(m *pb.Message) Message {
	msg := Message{
		ID:             uuid.MustParse(m.Id),
		ConversationID: uuid.MustParse(m.ConversationId),
		Role:           Role(m.Role),
		Content:        MessageContentFromPb(m.Content),
		CreatedAt:      m.CreatedAt.AsTime(),
	}

	if m.UserId != "" {
		uid := uuid.MustParse(m.UserId)
		msg.User = &User{ID: uid}
	}

	return msg
}

// ProviderFromPb converts a protobuf Provider to domain Provider
func ProviderFromPb(p *pb.Provider) Provider {
	return Provider{
		ID:        uuid.MustParse(p.Id),
		Name:      p.Name,
		CreatedAt: p.CreatedAt.AsTime(),
		UpdatedAt: p.UpdatedAt.AsTime(),
	}
}

// ModelFromPb converts a protobuf Model to domain Model
func ModelFromPb(m *pb.Model) Model {
	return Model{
		ID:          uuid.MustParse(m.Id),
		Provider:    Provider{ID: uuid.MustParse(m.ProviderId)}, // Partial provider
		Name:        m.Name,
		Description: m.Description,
		CreatedAt:   m.CreatedAt.AsTime(),
	}
}

// AIConfigFromPb converts a protobuf AIConfig to domain AIConfig
func AIConfigFromPb(a *pb.AIConfig) AIConfig {
	config := AIConfig{
		ID:           uuid.MustParse(a.Id),
		Name:         a.Name,
		Model:        ModelFromPb(a.Model),
		SystemPrompt: a.SystemPrompt,
		CreatedAt:    a.CreatedAt.AsTime(),
		UpdatedAt:    a.UpdatedAt.AsTime(),
	}

	if a.LastUsedAt != nil {
		t := a.LastUsedAt.AsTime()
		config.LastUsedAt = &t
	}

	return config
}
