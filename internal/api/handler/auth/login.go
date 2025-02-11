package auth

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
	"github.com/Dnlbb/auth/pkg/auth_v1"
)

// Login аутентификация пользователя.
func (c Controller) Login(ctx context.Context, request *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	user := models.User{Name: request.Username, Password: request.Password}

	refreshToken, err := c.authorizationService.Login(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("login failed: %w", err)
	}

	return &auth_v1.LoginResponse{RefreshToken: *refreshToken}, nil
}
