package repointerface

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

// CacheInterface интерфейс для работы с кэшем.
type CacheInterface interface {
	Create(ctx context.Context, id int64, user models.User) error
	Get(ctx context.Context, params models.GetUserParams) (*models.User, error)
}

// AccessPolicies политики доступа для юзеров (доступ определенных ролей х эндпоинтам).
type AccessPolicies interface {
	Check(path string, role string) error
}
