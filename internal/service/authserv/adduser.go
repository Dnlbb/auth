package authserv

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
	pgmodels "github.com/Dnlbb/auth/internal/repository/postgres/models"
)

func (s service) AddUser(ctx context.Context, user models.UserAdd) (*int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.storage.Save(ctx, user)
		if errTx != nil {
			return errTx
		}

		errTx = s.storage.Log(ctx, pgmodels.SAVE)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error when saving the user: %w", err)
	}
	return &id, nil
}
