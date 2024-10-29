package auth

import (
	"context"
	"fmt"

	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
)

// Create конвертация grpc структуры в сервисную модель и дальнейшая передача запроса в сервисный слой Create.
func (c *Controller) Create(ctx context.Context, req *authv1.CreateRequest) (*authv1.CreateResponse, error) {
	user := toModelUser(req)

	id, err := c.authService.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error while creating: %w", err)
	}

	return &authv1.CreateResponse{Id: *id}, nil
}
