package domain

import (
	pb "github.com/curator4/io/backend/internal/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MessageContentToPb converts domain MessageContent to protobuf MessageContent
func MessageContentToPb(c MessageContent) *pb.MessageContent {
	content := &pb.MessageContent{
		Text: c.Text,
	}

	// Convert media items
	if len(c.Media) > 0 {
		content.Media = make([]*pb.MediaItem, len(c.Media))
		for i, item := range c.Media {
			content.Media[i] = &pb.MediaItem{
				Type:     item.Type,
				Url:      item.URL,
				FileName: item.FileName,
			}
		}
	}

	return content
}

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
	msg := &pb.Message{
		Id:             m.ID.String(),
		ConversationId: m.ConversationID.String(),
		Role:           string(m.Role),
		Content:        MessageContentToPb(m.Content),
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
		ProviderId:  m.Provider.ID.String(),
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

// ActionToPb converts a domain DiscordAction to protobuf Action
func ActionToPb(a DiscordAction) *pb.Action {
	action := &pb.Action{}

	switch a.Type {
	case ActionTypeReaction:
		if a.Reaction != nil {
			action.ActionType = &pb.Action_Reaction{
				Reaction: &pb.ReactionAction{
					Emoji: a.Reaction.Emoji,
				},
			}
		}

	case ActionTypeWebSearch:
		if a.WebSearchResult != nil {
			// Convert web search items
			items := make([]*pb.WebSearchItem, len(a.WebSearchResult.Results))
			for i, item := range a.WebSearchResult.Results {
				items[i] = &pb.WebSearchItem{
					Title:   item.Title,
					Url:     item.URL,
					Snippet: item.Snippet,
				}
			}

			action.ActionType = &pb.Action_WebSearch{
				WebSearch: &pb.WebSearchAction{
					Queries: a.WebSearchResult.Queries,
					Results: items,
				},
			}
		}

	case ActionTypeImageGeneration:
		if a.ImageGeneration != nil {
			action.ActionType = &pb.Action_ImageGeneration{
				ImageGeneration: &pb.ImageGenerationAction{
					ImageData: a.ImageGeneration.ImageURL,
				},
			}
		}

	case ActionTypeCodeInterpreter:
		if a.CodeInterpreter != nil {
			action.ActionType = &pb.Action_CodeInterpreter{
				CodeInterpreter: &pb.CodeInterpreterAction{
					Code:    a.CodeInterpreter.Code,
					Outputs: a.CodeInterpreter.Outputs,
				},
			}
		}
	}

	return action
}

// ConversationLifecycleToPb converts domain ConversationLifecycle to protobuf
func ConversationLifecycleToPb(lc ConversationLifecycle) *pb.ConversationLifecycle {
	return &pb.ConversationLifecycle{
		IsNewConversation: lc.IsNewConversation,
		ConversationId:    lc.ConversationID.String(),
		ConversationName:  lc.ConversationName,
		StartedAt:         timestamppb.New(lc.StartedAt),
	}
}
