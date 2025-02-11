package redis

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Dnlbb/auth/internal/models"
	redisModels "github.com/Dnlbb/auth/internal/repository/models"
	"github.com/gomodule/redigo/redis"
)

func (c cache) GetById(ctx context.Context, id int) (*models.User, error) {
	key := strconv.Itoa(id)

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
