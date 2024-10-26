package pgconverter

import (
	"github.com/Dnlbb/auth/internal/models"
	"github.com/Dnlbb/auth/internal/repository/postgres/models"
)

// Repo2ServiceUser конвертируем user из модели, с которой работает postgresql в модель, с которой будет работать service.
func Repo2ServiceUser(pguser pgmodels.User) models.User {
	return models.User{ID: pguser.ID,
		Email:     pguser.Email,
		Password:  pguser.Password,
		Name:      pguser.Name,
		Role:      pguser.Role,
		CreatedAt: pguser.CreatedAt,
		UpdatedAt: pguser.UpdatedAt}
}
