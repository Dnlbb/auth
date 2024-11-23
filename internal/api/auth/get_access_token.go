package auth

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/pkg/auth_v1"
)

// GetAccessToken получение access токена по refresh токену.
func (c Controller) GetAccessToken(ctx context.Context, request *auth_v1.GetAccessTokenRequest) (*auth_v1.GetAccessTokenResponse, error) {
	accessToken, err := c.authorizationService.GetAccessToken(ctx, request.GetRefreshToken())
	if err != nil {
		return nil, fmt.Errorf("error getting access token: %w", err)
	}

	return &auth_v1.GetAccessTokenResponse{AccessToken: *accessToken}, nil
}
