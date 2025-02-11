package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
)

func (s service) GetById(ctx context.Context, id int) (*models.User, error) {
	var (
		userProfile *models.User
		errCache    error
		err         error
	)

	userProfile, errCache = s.cache.GetById(ctx, id)
	if errCache != nil {
		if errors.Is(errCache, models.ErrUserNotFound) {
			err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
				var errTx error
				userProfile, errTx = s.storage.GetUserById(ctx, id)
				if errTx != nil {
					return fmt.Errorf("error getting user profile: %w", errTx)
				}

				errTx = s.storage.Log(ctx, models.GETUSER)
				if errTx != nil {
					return fmt.Errorf("error logging: %w", errTx)
				}

				if errTx = s.cache.Create(ctx, userProfile.ID, *userProfile); errTx != nil {
					return fmt.Errorf("error caching user profile: %w", errTx)
				}

				return nil
			})

			if err != nil {
				return nil, err
			}
		}

		return nil, fmt.Errorf("error with cache: %w", errCache)
	}

	return userProfile, nil
}
