package auth

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/models"
	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *authv1.DeleteRequest) (*emptypb.Empty, error) {
	idDel := models.DeleteID(req.GetId())
	err := i.authService.DeleteUser(ctx, idDel)
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("error deleting user: %w", err)
	}

	return &emptypb.Empty{}, nil
}
