package storage

import (
	"context"
	"fmt"
	"log"

	"github.com/Dnlbb/auth/internal/client/db"
	"github.com/Dnlbb/auth/internal/models"
	pgmodels "github.com/Dnlbb/auth/internal/repository/postgres/models"
)

// Save сохранение пользователя в базу postgresql.
func (s *storage) Save(ctx context.Context, user models.UserAdd) (int64, error) {
	var userID pgmodels.UserID
	q := db.Query{
		Name:     "Save User",
		QueryRow: "INSERT INTO users (name, email, role, password) VALUES ($1, $2, $3, $4) RETURNING id",
	}
	err := s.db.DB().ScanOneContext(ctx, &userID, q, user.Name, user.Email, user.Role, user.Password)
	if err != nil {
		return 0, fmt.Errorf("error inserting user into database: %w", err)
	}
	log.Printf("Inserted user ID: %d", userID)
	return userID.ID, nil
}
