package redis

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Dnlbb/auth/internal/models"
)

func (c cache) Create(ctx context.Context, id int64, user models.User) error {
	idFormatted := strconv.FormatInt(id, 10)

	redisUser := toRedisModels(id, user)
	if err := c.cl.HashSet(ctx, idFormatted, redisUser); err != nil {
		return fmt.Errorf("failed to hash user: %w", err)
	}

	if err := c.cl.HashSet(ctx, user.Name, redisUser); err != nil {
		return fmt.Errorf("failed to hash user: %w", err)
	}

	if err := c.cl.Expire(ctx, idFormatted, 5*time.Minute); err != nil {
		return fmt.Errorf("failed to set expiration for user: %w", err)
	}
	return nil
}
