package authserv

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
)

func (s service) Get(ctx context.Context, params models.GetUserParams) (*models.User, error) {
	var userProfile *models.User

	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		userProfile, errTx = s.storage.GetUser(ctx, params)
		if errTx != nil {
			return errTx
		}

		errTx = s.storage.Log(ctx, models.GETUSER)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error while getting the user profile: %w", err)
	}

	return userProfile, nil
}
