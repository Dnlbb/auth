package authserv

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
	pgmodels "github.com/Dnlbb/auth/internal/repository/postgres/models"
)

func (s service) UpdateUser(ctx context.Context, userUpdate models.UpdateUser) error {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.storage.Update(ctx, userUpdate)
		if errTx != nil {
			return errTx
		}

		errTx = s.storage.Log(ctx, pgmodels.UPDATE)
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
