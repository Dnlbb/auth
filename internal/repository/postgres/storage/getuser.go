package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
	"github.com/Dnlbb/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

// GetUser получаем пользователя из базы postgresql
func (s *storage) GetUser(ctx context.Context, params models.GetUserParams) (*models.User, error) {
	var (
		user models.User
		err  error
	)

	query := sq.Select("id",
		"name",
		"email",
		"role",
		"password",
		"created_at",
		"updated_at",
	).From("users")

	switch {
	case params.ID != nil:
		query = query.Where(sq.Eq{"id": *params.ID})
	case params.Username != nil:
		query = query.Where(sq.Eq{"name": *params.Username})
	}

	query = query.PlaceholderFormat(sq.Dollar)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building sql query: %w", err)
	}

	q := db.Query{
		Name:     "Get user",
		QueryRow: sqlQuery,
	}

	err = s.db.DB().ScanOneContext(ctx, &user, q, args...)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("error: the user does not exist %w", err)
	} else if err != nil {
		return nil, fmt.Errorf("error when accessing the database to get a user profile: %w", err)
	}

	return &user, nil
}
