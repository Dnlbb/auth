package user

import (
	"context"
	"fmt"

	userv1 "github.com/Dnlbb/auth/pkg/user_v1"
)

// Get конвертация grpc структуры в сервисную модель и дальнейшая передача запроса в сервисный слой Get.
func (c *Controller) Get(ctx context.Context, req *userv1.GetRequest) (*userv1.GetResponse, error) {

	userProfile, err := c.authService.Get(ctx, mappingUserParams(req))
	if err != nil {
		return nil, fmt.Errorf("error when getUser request: %w", err)
	}

	return toProtoUserProfile(*userProfile), nil
}
