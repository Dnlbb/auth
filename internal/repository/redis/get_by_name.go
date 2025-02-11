package redis

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
	redisModels "github.com/Dnlbb/auth/internal/repository/models"
	"github.com/gomodule/redigo/redis"
)

func (c cache) GetByName(ctx context.Context, name string) (*models.User, error) {
	key := name

	userCache, err := c.cl.HGetAll(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("error with get user cache: %w", err)
	}

	if userCache == nil {
		return nil, models.ErrUserNotFound
	}

	var userProfile redisModels.User
	err = redis.ScanStruct(userCache, &userProfile)
	if err != nil {
		return nil, fmt.Errorf("error scanning user profile: %w", err)
	}

	fmt.Printf("%v", userProfile)
	user, err := toServiceModels(userProfile)
	if err != nil {
		return nil, fmt.Errorf("error converting user profile: %w", err)
	}

	return user, nil
}
