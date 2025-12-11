package domain

import (
	pb "github.com/curator4/io/backend/internal/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// UserToPb converts a domain User to protobuf User
func UserToPb(u User) *pb.User {
	return &pb.User{
		Id:        u.ID.String(),
		Name:      u.Name,
		CreatedAt: timestamppb.New(u.CreatedAt),
		UpdatedAt: timestamppb.New(u.UpdatedAt),
	}
}

// ConversationToPb converts a domain Conversation to protobuf Conversation
func ConversationToPb(c Conversation) *pb.Conversation {
	return &pb.Conversation{
		Id:        c.ID.String(),
		Name:      c.Name,
		CreatedAt: timestamppb.New(c.CreatedAt),
		UpdatedAt: timestamppb.New(c.UpdatedAt),
	}
}

// MessageToPb converts a domain Message to protobuf Message
func MessageToPb(m Message) *pb.Message {
	content := &pb.MessageContent{
		Text: m.Content.Text,
	}

	// Convert media items
	if len(m.Content.Media) > 0 {
		content.Media = make([]*pb.MediaItem, len(m.Content.Media))
		for i, item := range m.Content.Media {
			content.Media[i] = &pb.MediaItem{
				Type:     item.Type,
				Url:      item.URL,
				FileName: item.FileName,
			}
		}
	}

	msg := &pb.Message{
		Id:             m.ID.String(),
		ConversationId: m.ConversationID.String(),
		Role:           string(m.Role),
		Content:        content,
		CreatedAt:      timestamppb.New(m.CreatedAt),
	}

	if m.User != nil {
		msg.UserId = m.User.ID.String()
	}

	return msg
}

// ProviderToPb converts a domain Provider to protobuf Provider
func ProviderToPb(p Provider) *pb.Provider {
	return &pb.Provider{
		Id:        p.ID.String(),
		Name:      p.Name,
		CreatedAt: timestamppb.New(p.CreatedAt),
		UpdatedAt: timestamppb.New(p.UpdatedAt),
	}
}

// ModelToPb converts a domain Model to protobuf Model
func ModelToPb(m Model) *pb.Model {
	return &pb.Model{
		Id:          m.ID.String(),
		ProviderId:  m.ProviderID.String(),
		Name:        m.Name,
		Description: m.Description,
		CreatedAt:   timestamppb.New(m.CreatedAt),
	}
}

// AIConfigToPb converts a domain AIConfig to protobuf AIConfig
func AIConfigToPb(a AIConfig) *pb.AIConfig {
	config := &pb.AIConfig{
		Id:           a.ID.String(),
		Name:         a.Name,
		Model:        ModelToPb(a.Model),
		SystemPrompt: a.SystemPrompt,
		CreatedAt:    timestamppb.New(a.CreatedAt),
		UpdatedAt:    timestamppb.New(a.UpdatedAt),
	}

	if a.LastUsedAt != nil {
		config.LastUsedAt = timestamppb.New(*a.LastUsedAt)
	}

	return config
}
