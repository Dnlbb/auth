package authserv

import (
	"context"
	"errors"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
	"github.com/Dnlbb/auth/internal/repository/redis"
)

func (s service) Get(ctx context.Context, params models.GetUserParams) (*models.User, error) {
	var (
		userProfile *models.User
		errCache    error
		err         error
	)

	userProfile, errCache = s.cache.Get(ctx, params)
	if errors.Is(errCache, redis.ErrUserNotFound) {
		err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
			var errTx error
			userProfile, errTx = s.storage.GetUser(ctx, params)
			if errTx != nil {
				return fmt.Errorf("error getting user profile: %w", errTx)
			}

			errTx = s.storage.Log(ctx, models.GETUSER)
			if errTx != nil {
				return fmt.Errorf("error logging: %w", errTx)
			}

			if errTx = s.cache.Create(ctx, userProfile.ID, *userProfile); err != nil {
				return fmt.Errorf("error caching user profile: %w", errTx)
			}

			return nil
		})

		if err != nil {
			return nil, err
		}

		return userProfile, nil

	} else if errCache != nil {
		return nil, fmt.Errorf("error with cache: %w", errCache)
	}

	return userProfile, nil
}
