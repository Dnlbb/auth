package auth

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete конвертация grpc структуры в сервисную модель и дальнейшая передача запроса в сервисный слой Delete.
func (c *Controller) Delete(ctx context.Context, req *authv1.DeleteRequest) (*emptypb.Empty, error) {
	idDel := models.DeleteID(req.GetId())

	if err := c.authService.Delete(ctx, idDel); err != nil {
		return &emptypb.Empty{}, fmt.Errorf("deleting user error: %w", err)
	}

	return &emptypb.Empty{}, nil
}
