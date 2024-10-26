package servinterfaces

import (
	"context"

	"github.com/Dnlbb/auth/internal/models"
)

// AuthService интерфейс сервиса
type AuthService interface {
	AddUser(ctx context.Context, user models.UserAdd) (*int64, error)
	DeleteUser(ctx context.Context, userID models.DeleteID) error
	UpdateUser(ctx context.Context, userUpdate models.UpdateUser) error
	GetUser(ctx context.Context, params models.GetUserParams) (*models.User, error)
}
