package auth

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/converter"
	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *authv1.UpdateRequest) (*emptypb.Empty, error) {
	updateUser := converter.ProtoUpdateUser2UpdateUser(req)
	err := i.authService.UpdateUser(ctx, *updateUser)
	if err != nil {
		return &emptypb.Empty{}, fmt.Errorf("error updating user: %w", err)
	}
	return &emptypb.Empty{}, nil
}
