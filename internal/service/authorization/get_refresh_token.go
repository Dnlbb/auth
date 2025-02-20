package authorization

import (
	"context"

	"github.com/Dnlbb/auth/internal/models"
	"github.com/Dnlbb/auth/internal/service/authorization/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s service) GetRefreshToken(ctx context.Context, refreshToken string) (*string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(s.config.GetRefreshTokenSecretKey()))
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "invalid refresh token")
	}

	user, err := s.storage.GetUser(ctx, models.GetUserParams{Username: &claims.Username})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	timeExpiration, err := s.config.GetRefreshTokenExpiration()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "could not get refresh token expiration")
	}

	accessToken, err := utils.GenerateToken(models.UserPayload{
		Username: user.Name,
		Role:     user.Role,
	},
		[]byte(s.config.GetAccessTokenSecretKey()),
		timeExpiration,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "unable to generate access token")
	}

	return &accessToken, nil
}
