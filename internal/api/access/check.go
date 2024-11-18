package access

import (
	"context"

	"github.com/Dnlbb/auth/pkg/access_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Check передача параметров в сервисный слой.
func (c Controller) Check(ctx context.Context, request *access_v1.CheckRequest) (*emptypb.Empty, error) {
	if access := c.accessService.Check(ctx, request.GetEndpointAddress()); access != nil {
		return &emptypb.Empty{}, access
	}

	return &emptypb.Empty{}, nil
}
