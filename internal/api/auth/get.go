package auth

import (
	"context"
	"fmt"

	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
)

// Get конвертация grpc структуры в сервисную модель и дальнейшая передача запроса в сервисный слой Get.
func (c *Controller) Get(ctx context.Context, req *authv1.GetRequest) (*authv1.GetResponse, error) {
	params := mappingUserParams(req)

	userProfile, err := c.authService.Get(ctx, *params)
	if err != nil {
		return nil, fmt.Errorf("error when getUser request: %w", err)
	}

	return toProtoUserProfile(*userProfile), nil
}
