package core

import (
	"context"
	"fmt"

	"github.com/curator4/io/backend/internal/domain"
	"github.com/google/uuid"
)

// getActiveConfig gets the active ai config or loads default
func (c *Core) getActiveConfig(ctx context.Context) (*domain.AIConfig, error) {
	c.mu.RLock()
	config := c.session.ActiveConfig
	c.mu.RUnlock()

	// check if activeconfig exists and return if it does
	if config != nil {
		return config, nil
	}

	// no active ai config, load from db
	configs, err := c.db.ListAIConfigs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list configs: %w", err)
	}
	if len(configs) == 0 {
		return nil, fmt.Errorf("no ai configs in db")
	}

	// load first config
	configRow, err := c.db.GetAIConfigByID(ctx, configs[0].ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get first config: %w", err)
	}

	defaultConfig := domain.AIConfigFromDB(configRow)

	// set activeConfig
	c.mu.Lock()
	c.session.ActiveConfig = &defaultConfig
	c.mu.Unlock()

	return &defaultConfig, nil
}

// SetActiveConfig sets the active ai config
func (c *Core) SetActiveConfig(ctx context.Context, configID uuid.UUID) error {
	configRow, err := c.db.GetAIConfigByID(ctx, configID)
	if err != nil {
		return fmt.Errorf("failed to get config: %w", err)
	}

	config := domain.AIConfigFromDB(configRow)

	c.mu.Lock()
	c.session.ActiveConfig = &config
	c.mu.Unlock()

	return nil
}
