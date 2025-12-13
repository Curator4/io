package core

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/curator4/io/backend/internal/database"
	"github.com/curator4/io/backend/internal/domain"
	"github.com/google/uuid"
)

// getOrCreateUser gets a domain.User by name from the database. if username doesn't exist, creates it
func (c *Core) getOrCreateUser(ctx context.Context, username string) (domain.User, error) {
	// normalize username to lowercase
	username = strings.ToLower(username)

	// check if it exists in database
	dbUser, err := c.db.GetUserByName(ctx, username)

	// if it doesn't, create new user
	if err == sql.ErrNoRows {
		params := database.CreateUserParams{
			ID:   uuid.New(),
			Name: username,
		}

		dbUser, err = c.db.CreateUser(ctx, params)

		if err != nil {
			return domain.User{}, fmt.Errorf("failed to create user: %w", err)
		}
	} else if err != nil {
		return domain.User{}, fmt.Errorf("failed to get user: %w", err)
	}

	return domain.UserFromDB(dbUser), nil
}
