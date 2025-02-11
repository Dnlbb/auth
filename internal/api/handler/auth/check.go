package auth

import (
	"context"

	"github.com/Dnlbb/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Check проверка доступа.
func (c Controller) Check(ctx context.Context, request *auth_v1.CheckRequest) (*emptypb.Empty, error) {
	if access := c.authorizationService.Check(ctx, request.GetEndpointAddress()); access != nil {
		return &emptypb.Empty{}, access
	}

	return &emptypb.Empty{}, nil
}
