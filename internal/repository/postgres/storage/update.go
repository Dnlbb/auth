package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/Dnlbb/auth/internal/client/db"
	"github.com/Dnlbb/auth/internal/models"
)

// Update обновления данных пользователя в базе postgresql.
func (s *storage) Update(ctx context.Context, update models.UpdateUser) error {
	q := db.Query{
		Name:     "Update User",
		QueryRow: "UPDATE users SET name = $1, email = $2, role = $3, password = $4 WHERE id = $5",
	}
	res, err := s.db.DB().ExecContext(ctx, q, update.Name, update.Email, update.Role, update.Password, update.ID)
	if err != nil {
		return fmt.Errorf("error updating user into database: %w", err)
	}
	log.Printf("Updated user: %v", res.RowsAffected())
	return nil
}
