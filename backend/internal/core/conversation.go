package core

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/curator4/io/backend/internal/database"
	"github.com/curator4/io/backend/internal/domain"
	"github.com/google/uuid"
)

// createConversation creates a new conversation, given a name
func (c *Core) createConversation(ctx context.Context, name string) (domain.Conversation, error) {
	dbConv, err := c.db.CreateConversation(ctx, sql.NullString{
		String: name,
		Valid:  true,
	})
	if err != nil {
		return domain.Conversation{}, fmt.Errorf("failed to create conversation %w", err)
	}
	return domain.ConversationFromDB(dbConv), nil
}

// getConversationHistory gets all the messages from a conversation, given a conversation ID
func (c *Core) getConversationHistory(ctx context.Context, conversationID uuid.UUID) ([]domain.Message, error) {
	rows, err := c.db.GetMessagesByConversation(ctx, conversationID)
	if err != nil {
		return nil, fmt.Errorf("failed to get conversation history %w", err)
	}

	messages := make([]domain.Message, len(rows))
	for i, row := range rows {
		messages[i] = domain.MessageFromDB(row)
	}

	return messages, nil
}

// getOrCreateActiveConversation returns or creates and returns the active conversation
func (c *Core) getOrCreateActiveConversation(ctx context.Context) (*domain.Conversation, error) {
	c.mu.RLock()
	activeConv := c.session.ActiveConversation
	lastActivity := c.session.LastActivity
	c.mu.RUnlock()

	// check if there is active conversation (checks < 30 min)
	if activeConv != nil && time.Since(lastActivity) < 30*time.Minute {
		return activeConv, nil
	}

	// create new conversation
	conversationName := time.Now().Format("Jan 2, 2006 15:04")
	conv, err := c.createConversation(ctx, conversationName)
	if err != nil {
		return nil, err
	}

	// set as active
	c.mu.Lock()
	c.session.ActiveConversation = &conv
	c.mu.Unlock()

	// update last_used_at in db
	if err := c.db.UpdateConversationLastUsed(ctx, conv.ID); err != nil {
		return nil, fmt.Errorf("failed to updated conversation last used: %w", err)
	}

	return &conv, nil
}

// addParticipantIfNeeded
func (c *Core) addParticipantIfNeeded(ctx context.Context, conv *domain.Conversation, user domain.User) error {
	// check if participant
	for _, p := range conv.Participants {
		if p.UserID == user.ID {
			return nil
		}
	}

	// add to database
	params := database.AddParticipantParams{
		ConversationID: conv.ID,
		UserID:         user.ID,
	}
	err := c.db.AddParticipant(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to add conversation participant: %w", err)
	}

	// add to memory
	participant := domain.ConversationParticipant{
		ConversationID: conv.ID,
		UserID:         user.ID,
		JoinedAt:       time.Now(),
	}
	conv.Participants = append(conv.Participants, participant)

	return nil
}
