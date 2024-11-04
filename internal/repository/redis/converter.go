package redis

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Dnlbb/auth/internal/models"
	redismodels "github.com/Dnlbb/auth/internal/repository/models"
)

func toRedisModels(id int64, user models.User) redismodels.User {
	idStr, timeNow := strconv.FormatInt(id, 10), time.Now()

	return redismodels.User{
		ID:        idStr,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: timeNow.Format(time.RFC3339),
		UpdatedAt: timeNow.Format(time.RFC3339),
	}
}

func toServiceModels(user redismodels.User) (*models.User, error) {
	CreatedAt, err := time.Parse(time.RFC3339, user.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("error with time parse CreatedAt")
	}
	UpdatedAt, err := time.Parse(time.RFC3339, user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error with time parse UpdatedAt")
	}
	id, err := strconv.ParseInt(user.ID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error with parse ID")
	}

	return &models.User{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}, nil
}
