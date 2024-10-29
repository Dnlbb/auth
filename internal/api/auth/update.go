package auth

import (
	"context"
	"fmt"

	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update конвертация grpc структуры в сервисную модель и дальнейшая передача запроса в сервисный слой Update.
func (c *Controller) Update(ctx context.Context, req *authv1.UpdateRequest) (*emptypb.Empty, error) {
	updateUser := toUpdateUser(req)

	if err := c.authService.Update(ctx, *updateUser); err != nil {
		return &emptypb.Empty{}, fmt.Errorf("error updating user: %w", err)
	}

	return &emptypb.Empty{}, nil
}
