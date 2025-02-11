package auth

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/pkg/auth_v1"
)

// GetRefreshToken получения нового refresh токена по старому.
func (c Controller) GetRefreshToken(ctx context.Context, request *auth_v1.GetRefreshTokenRequest) (*auth_v1.GetRefreshTokenResponse, error) {
	refreshToken, err := c.authorizationService.GetRefreshToken(ctx, request.GetOldRefreshToken())
	if err != nil {
		return nil, fmt.Errorf("could not get refresh token: %w", err)
	}

	return &auth_v1.GetRefreshTokenResponse{RefreshToken: *refreshToken}, nil
}
