package authserv

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
)

func (s service) Create(ctx context.Context, user models.User) (*int64, error) {
	var id int64

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

	return &id, nil
}
