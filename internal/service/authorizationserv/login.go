package authorizationserv

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
	"github.com/Dnlbb/auth/internal/service/authorizationserv/utils"
)

func (s service) Login(ctx context.Context, user models.User) (*string, error) {
	User, err := s.storage.GetUser(ctx, models.GetUserParams{Username: &user.Name})
	if err != nil {
		return nil, fmt.Errorf("get user: %w", err)
	}

	if User.Password != user.Password {
		return nil, fmt.Errorf("invalid password")
	}

	timeExpiration, err := s.config.GetRefreshTokenExpiration()
	if err != nil {
		return nil, fmt.Errorf("get refresh token secret key: %w", err)
	}

	refreshToken, err := utils.GenerateToken(models.UserPayload{
		Username: user.Name,
		Role:     User.Role,
	},
		[]byte(s.config.GetRefreshTokenSecretKey()),
		timeExpiration,
	)

	if err != nil {
		return nil, fmt.Errorf("error generate refresh token: %w", err)
	}

	return &refreshToken, nil
}
