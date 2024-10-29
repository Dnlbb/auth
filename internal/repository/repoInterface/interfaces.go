package repoInterface

import (
	"context"

	"github.com/Dnlbb/auth/internal/models"
)

// StorageInterface интерфейс для работы с хранилищем.
type StorageInterface interface {
	Save(ctx context.Context, user models.User) (int64, error)
	Update(ctx context.Context, update models.User) error
	Delete(ctx context.Context, id models.DeleteID) error
	GetUser(ctx context.Context, params models.GetUserParams) (*models.User, error)
	Log(ctx context.Context, key models.LogKey) error
}
