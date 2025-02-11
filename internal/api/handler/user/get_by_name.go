package user

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/api/request"
	userv1 "github.com/Dnlbb/auth/pkg/user_v1"
)

// GetByName конвертация grpc структуры в сервисную модель и дальнейшая передача запроса в сервисный слой Get.
func (c *Controller) GetByName(ctx context.Context, req *userv1.GetByNameRequest) (*userv1.GetByResponse, error) {
	name, err := request.Name(req)
	if err != nil {
		return nil, err
	}

	userProfile, err := c.authService.GetByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("error when getUser request: %w", err)
	}

	return toProtoUserProfile(*userProfile), nil
}
