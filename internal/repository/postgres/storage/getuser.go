package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/Dnlbb/auth/internal/client/db"
	"github.com/Dnlbb/auth/internal/models"
	pgconverter "github.com/Dnlbb/auth/internal/repository/postgres/converter"
	pgmodels "github.com/Dnlbb/auth/internal/repository/postgres/models"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

// GetUser получаем пользователя из базы postgresql
func (s *storage) GetUser(ctx context.Context, params models.GetUserParams) (*models.User, error) {
	var pgUser pgmodels.User
	var err error
	query := sq.Select(IDColumn, UsernameColumn, EmailColumn, RoleColumn, CreateTimeColumn, UpdateTimeColumn).From(UserTableName).PlaceholderFormat(sq.Dollar)
	switch {
	case params.ID != nil:
		query = query.Where(sq.Eq{IDColumn: *params.ID})
	case params.Username != nil:
		query = query.Where(sq.Eq{UsernameColumn: *params.Username})
	default:
		return nil, fmt.Errorf("error: ID and USERNAME is null")
	}
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building sql query: %w", err)
	}
	q := db.Query{
		Name:     "Get user",
		QueryRow: sqlQuery,
	}

	err = s.db.DB().ScanOneContext(ctx, &pgUser, q, args...)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("пользователя не существует")
	} else if err != nil {
		return nil, fmt.Errorf("ошибка при обращении в базу для получения профиля пользователя: %v", err)
	}
	user := pgconverter.Repo2ServiceUser(pgUser)
	return &user, nil
}
