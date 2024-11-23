package authorization

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/Dnlbb/auth/internal/service/authorization/utils"
	"google.golang.org/grpc/metadata"
)

// Check проверка доступа.
func (s service) Check(ctx context.Context, address string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("metadata is not provided")
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return errors.New("authorization header is not provided")
	}

	if !strings.HasPrefix(authHeader[0], "Bearer ") {
		return errors.New("invalid authorization header format")
	}

	accessToken := strings.TrimPrefix(authHeader[0], "Bearer ")

	claims, err := utils.VerifyToken(accessToken, []byte(s.config.GetAccessTokenSecretKey()))
	if err != nil {
		return errors.New("access token is invalid")
	}

	if access := s.accessPolicy.Check(address, claims.Role); access != nil {
		return fmt.Errorf("access policy is not allowed: %w", access)
	}

	return nil
}
