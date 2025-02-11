package user

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/api/request"
	userv1 "github.com/Dnlbb/auth/pkg/user_v1"
)

// GetById конвертация grpc структуры в сервисную модель и дальнейшая передача запроса в сервисный слой Get.
func (c *Controller) GetById(ctx context.Context, req *userv1.GetByIdRequest) (*userv1.GetByResponse, error) {
	id := request.ID(req)

	userProfile, err := c.authService.GetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error when getUser request: %w", err)
	}

	return toProtoUserProfile(*userProfile), nil
}
