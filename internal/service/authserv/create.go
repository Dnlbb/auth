package authserv

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
)

func (s service) Create(ctx context.Context, user models.User) (*int64, error) {
	var id int64
	if err := validatePass(user.Password); err != nil {
		return nil, fmt.Errorf("invalid password: %w", err)
	}

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.storage.Save(ctx, user)
		if errTx != nil {
			return errTx
		}

		errTx = s.storage.Log(ctx, models.SAVE)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("user creation error: %w", err)
	}

	if err = s.cache.Create(ctx, id, user); err != nil {
		return nil, fmt.Errorf("user add cache error: %w", err)
	}

	return &id, nil
}

func validatePass(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters")
	} else if len(password) > 255 {
		return fmt.Errorf("password must be at most 255 characters")
	}
	return nil
}
