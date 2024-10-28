package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/Dnlbb/auth/internal/client/db"
	"github.com/Dnlbb/auth/internal/models"
)

// Delete удаления пользователя из базы postgresql.
func (s *storage) Delete(ctx context.Context, id models.DeleteID) error {
	q := db.Query{
		Name:     "Delete user",
		QueryRow: "DELETE FROM users WHERE id = $1",
	}

	rows, err := s.db.DB().ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("error deleting user into database: %w", err)
	}

	if rowsAffected := rows.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("error: such user does not exist")
	}

	log.Printf("Deleted user: %v", id)

	return nil
}
