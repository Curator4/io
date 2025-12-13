package core

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/curator4/io/backend/internal/catalog"
	"github.com/curator4/io/backend/internal/database"
)

// syncCataloc syncs the providers/models defined in catalog package to database
// call on startup and after migrations
func syncCatalog(ctx context.Context, c *Core) error {
	// sync providers
	for _, provider := range catalog.GetAllProviders() {
		_, err := c.db.GetProvider(ctx, string(provider))
		if err == sql.ErrNoRows {
			_, err = c.db.CreateProvider(ctx, string(provider))
			if err != nil {
				return fmt.Errorf("failed to sync provider %s: %w", provider, err)
			}
		} else if err != nil {
			return fmt.Errorf("failed to check provider %s: %w", provider, err)
		}
	}

	// sync models
	for provider, models := range catalog.ValidModels {
		providerRow, err := c.db.GetProvider(ctx, string(provider))
		if err != nil {
			return fmt.Errorf("failed to get provider %s: %w", provider, err)
		}

		for _, model := range models {
			_, err := c.db.GetModelByName(ctx, string(model))
			if err == sql.ErrNoRows {
				params := database.CreateModelParams{
					ProviderID:  providerRow.ID,
					Name:        string(model),
					Description: sql.NullString{},
				}
				_, err = c.db.CreateModel(ctx, params)
				if err != nil {
					return fmt.Errorf("failed to sync model %s: %w", model, err)
				}
			} else if err != nil {
				return fmt.Errorf("failed to check model %s: %w", model, err)
			}
		}
	}

	return nil
}

// createDefaultAIConfig creates the default ai config using the values in catalog, sets as active
// only runs if aiconfig table is empty
func createDefaultAIConfig(ctx context.Context, c *Core) error {
	configs, err := c.db.ListAIConfigs(ctx)
	if err != nil {
		return fmt.Errorf("failed to list configs: %w", err)
	}

	if len(configs) == 0 {
		defaultID, err := c.CreateAIConfig(
			ctx,
			catalog.DefaultProviderName,
			catalog.DefaultModelName,
			catalog.DefaultConfigName,
			catalog.DefaultSystemPrompt,
		)

		if err != nil {
			return fmt.Errorf("failed to create default AI Config: %w", err)
		}

		if err = c.SetActiveConfig(ctx, defaultID); err != nil {
			return fmt.Errorf("failed to set active ai config in database init: %w", err)
		}
	}

	return nil
}

// InitializeDatabase syncs the database in accordance to the defined values in catalog package
// Adds both providers, models, and a default ai config
// doesn't remove values
func (c *Core) InitializeDatabase(ctx context.Context) error {
	if err := syncCatalog(ctx, c); err != nil {
		return err
	}
	if err := createDefaultAIConfig(ctx, c); err != nil {
		return err
	}

	return nil
}
