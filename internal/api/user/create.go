package user

import (
	"context"
	"fmt"

	userv1 "github.com/Dnlbb/auth/pkg/user_v1"
)

// Create конвертация grpc структуры в сервисную модель и дальнейшая передача запроса в сервисный слой Create.
func (c *Controller) Create(ctx context.Context, req *userv1.CreateRequest) (*userv1.CreateResponse, error) {
	user := toModelUser(req)

	id, err := c.authService.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error while creating: %w", err)
	}

	return &userv1.CreateResponse{Id: *id}, nil
}
