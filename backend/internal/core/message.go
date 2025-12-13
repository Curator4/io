package core

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/curator4/io/backend/internal/database"
	"github.com/curator4/io/backend/internal/domain"
	"github.com/google/uuid"
)

// storeMessage stores a message in the database
func (c *Core) storeMessage(
	ctx context.Context,
	conversationID uuid.UUID,
	user *domain.User,
	role domain.Role,
	content domain.MessageContent,
) (domain.Message, error) {

	// json marshall to store as jsonb
	contentJSON, err := json.Marshal(content)
	if err != nil {
		return domain.Message{}, fmt.Errorf("failed to marshal message content: %w", err)
	}

	var nullUserID uuid.NullUUID
	if user != nil {
		nullUserID = uuid.NullUUID{UUID: user.ID, Valid: true}
	}

	params := database.CreateMessageParams{
		ConversationID: conversationID,
		UserID:         nullUserID,
		Role:           string(role),
		Content:        contentJSON,
	}

	dbMsg, err := c.db.CreateMessage(ctx, params)
	if err != nil {
		return domain.Message{}, fmt.Errorf("failed to create message: %w", err)
	}

	// Construct from data we already have + DB-generated fields
	return domain.Message{
		ID:             dbMsg.ID,
		ConversationID: conversationID,
		User:           user,
		Role:           role,
		Content:        content,
		CreatedAt:      dbMsg.CreatedAt,
	}, nil
}
