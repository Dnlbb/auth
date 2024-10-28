package servinterfaces

import (
	"context"

	"github.com/Dnlbb/auth/internal/models"
)

// AuthService интерфейс сервиса
type AuthService interface {
	Create(ctx context.Context, user models.User) (*int64, error)
	Delete(ctx context.Context, userID models.DeleteID) error
	Update(ctx context.Context, userUpdate models.User) error
	Get(ctx context.Context, params models.GetUserParams) (*models.User, error)
}
