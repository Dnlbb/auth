package authserv

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
	pgmodels "github.com/Dnlbb/auth/internal/repository/postgres/models"
)

func (s service) GetUser(ctx context.Context, params models.GetUserParams) (*models.User, error) {
	var userProfile *models.User
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		userProfile, errTx = s.storage.GetUser(ctx, params)
		if errTx != nil {
			return errTx
		}

		errTx = s.storage.Log(ctx, pgmodels.GETUSER)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error when getting the user profile: %w", err)
	}

	return userProfile, nil
}
