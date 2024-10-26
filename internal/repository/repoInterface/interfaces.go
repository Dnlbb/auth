package repoInterface

import (
	"context"

	"github.com/Dnlbb/auth/internal/models"
	pgmodels "github.com/Dnlbb/auth/internal/repository/postgres/models"
)

type StorageInterface interface {
	Save(ctx context.Context, user models.UserAdd) (int64, error)
	Update(ctx context.Context, update models.UpdateUser) error
	Delete(ctx context.Context, id models.DeleteID) error
	GetUser(ctx context.Context, params models.GetUserParams) (*models.User, error)
	Log(ctx context.Context, key pgmodels.LogKey) error
}
