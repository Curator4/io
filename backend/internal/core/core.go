package core

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/curator4/io/backend/internal/database"
	"github.com/curator4/io/backend/internal/domain"
	"github.com/curator4/io/backend/internal/llm"
)

type Session struct {
	ActiveConfig       *domain.AIConfig
	ActiveConversation *domain.Conversation
	LastActivity       time.Time
}

type Core struct {
	db           *database.Queries
	llmProviders map[string]llm.Provider
	session      Session
	mu           sync.RWMutex
}

func NewCore(db *database.Queries, providers map[string]llm.Provider) *Core {
	return &Core{
		db:           db,
		llmProviders: providers,
		session:      Session{},
	}
}

// prepareAndStoreUserMessage handles common setup for both SendMessage and StoreMessage
func (c *Core) prepareAndStoreUserMessage(
	ctx context.Context,
	content domain.MessageContent,
	username string,
) (user domain.User, conv *domain.Conversation, err error) {
	// 1. get/create user
	user, err = c.getOrCreateUser(ctx, username)
	if err != nil {
		err = fmt.Errorf("failed to get or create user: %w", err)
		return
	}

	// 2. get/create conversation
	conv, err = c.getOrCreateActiveConversation(ctx)
	if err != nil {
		err = fmt.Errorf("failed to get or create active conversation: %w", err)
		return
	}

	// 3. add user as conversation participant
	if err = c.addParticipantIfNeeded(ctx, conv, user); err != nil {
		err = fmt.Errorf("failed to add participant: %w", err)
		return
	}

	// 4. store message in db
	_, err = c.storeMessage(ctx, conv.ID, &user, domain.RoleUser, content)
	if err != nil {
		err = fmt.Errorf("failed to store user message: %w", err)
		return
	}

	return
}

func (c *Core) HandleSendMessage(
	ctx context.Context,
	content domain.MessageContent,
	username string,
) (llmMsg domain.Message, err error) {

	// 1-4. prepare and store user message
	_, conv, err := c.prepareAndStoreUserMessage(ctx, content, username)
	if err != nil {
		return
	}

	// 5. get conversation history
	history, err := c.getConversationHistory(ctx, conv.ID)
	if err != nil {
		err = fmt.Errorf("failed to get conversation history: %w", err)
		return
	}

	// 6. get active ai config
	config, err := c.getActiveConfig(ctx)
	if err != nil {
		err = fmt.Errorf("failed to get active config: %w", err)
		return
	}

	// 7. get llm provider
	provider, ok := c.llmProviders[config.Model.Provider.Name]
	if !ok {
		err = fmt.Errorf("unknown provider: %s", config.Model.Provider.Name)
		return
	}

	// 8. call llm
	llmContent, err := provider.SendMessage(ctx, history, *config)
	if err != nil {
		err = fmt.Errorf("failed to call llm provider: %w", err)
		return
	}

	// 9. store assistant message
	llmMsg, err = c.storeMessage(ctx,
		conv.ID,
		nil,
		domain.RoleAssistant,
		llmContent,
	)
	if err != nil {
		err = fmt.Errorf("failed to store assistant message: %w", err)
		return
	}

	// 10. updated session
	c.mu.Lock()
	c.session.LastActivity = time.Now()
	c.mu.Unlock()

	return
}

func (c *Core) HandleStoreMessage(
	ctx context.Context,
	content domain.MessageContent,
	username string,
) (err error) {

	// 1-4. prepare and store user message
	_, _, err = c.prepareAndStoreUserMessage(ctx, content, username)
	if err != nil {
		return
	}

	// 5. update session
	c.mu.Lock()
	c.session.LastActivity = time.Now()
	c.mu.Unlock()

	return
}
