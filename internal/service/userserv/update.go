package userserv

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
)

func (s service) Update(ctx context.Context, userUpdate models.User) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.storage.Update(ctx, userUpdate)
		if errTx != nil {
			return errTx
		}

		errTx = s.storage.Log(ctx, models.UPDATE)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("could not update user: %w", err)
	}

	return nil
}
