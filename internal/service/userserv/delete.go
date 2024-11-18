package userserv

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
)

func (s service) Delete(ctx context.Context, userID models.DeleteID) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.storage.Delete(ctx, userID)
		if errTx != nil {
			return errTx
		}

		errTx = s.storage.Log(ctx, models.DELETE)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}
