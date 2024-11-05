package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/Dnlbb/auth/internal/models"
	"github.com/Dnlbb/platform_common/pkg/db"
)

// Save сохранение пользователя в базу postgresql.
func (s *storage) Save(ctx context.Context, user models.User) (int64, error) {
	var userID int64
	q := db.Query{
		Name:     "Save User",
		QueryRow: "INSERT INTO users (name, email, role, password) VALUES ($1, $2, $3, $4) RETURNING id",
	}

	err := s.db.DB().ScanOneContext(ctx, &userID, q, user.Name, user.Email, user.Role, user.Password)
	if err != nil {
		return 0, fmt.Errorf("error inserting user into database: %w", err)
	}

	log.Printf("Inserted user ID: %d", userID)

	return userID, nil
}
