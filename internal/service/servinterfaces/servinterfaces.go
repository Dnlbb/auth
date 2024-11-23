package servinterfaces

import (
	"context"

	"github.com/Dnlbb/auth/internal/models"
)

// UserService интерфейс сервиса.
type UserService interface {
	Create(ctx context.Context, user models.User) (*int64, error)
	Delete(ctx context.Context, userID models.DeleteID) error
	Update(ctx context.Context, userUpdate models.User) error
	Get(ctx context.Context, params models.GetUserParams) (*models.User, error)
}

// AuthorizationService интерфейс реализации авторизации и аутентификации.
type AuthorizationService interface {
	Login(ctx context.Context, user models.User) (*string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (*string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (*string, error)
	Check(ctx context.Context, address string) error
}
