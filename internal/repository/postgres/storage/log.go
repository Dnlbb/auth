package storage

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
	"github.com/Dnlbb/platform_common/pkg/db"
	sq "github.com/Masterminds/squirrel"
)

func (s *storage) Log(ctx context.Context, key models.LogKey) error {
	query := sq.Insert("log").Columns("name")

	switch key {
	case models.SAVE:
		query = query.Values(models.SAVE)
	case models.GETUSER:
		query = query.Values(models.GETUSER)
	case models.DELETE:
		query = query.Values(models.DELETE)
	case models.UPDATE:
		query = query.Values(models.UPDATE)
	}

	query = query.PlaceholderFormat(sq.Dollar)

	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("error with make sql query: %w", err)
	}

	q := db.Query{
		Name:     "Log",
		QueryRow: sql,
	}

	_, err = s.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return fmt.Errorf("error with insert to logging table: %w", err)
	}

	return nil
}
