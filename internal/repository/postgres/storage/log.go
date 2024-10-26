package storage

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/client/db"
	pgmodels "github.com/Dnlbb/auth/internal/repository/postgres/models"
	sq "github.com/Masterminds/squirrel"
)

func (s *storage) Log(ctx context.Context, key pgmodels.LogKey) error {
	query := sq.Insert(UserLogName).Columns(QueryName)

	switch key {
	case pgmodels.SAVE:
		query = query.Values(pgmodels.SAVE)
	case pgmodels.GETUSER:
		query = query.Values(pgmodels.GETUSER)
	case pgmodels.DELETE:
		query = query.Values(pgmodels.DELETE)
	case pgmodels.UPDATE:
		query = query.Values(pgmodels.UPDATE)
	default:
		return fmt.Errorf("invalid LogKey: %v", key)
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
