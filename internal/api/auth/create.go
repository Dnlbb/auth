package auth

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/converter"
	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
)

// Create реализация сгенерированного grpc
func (i *Implementation) Create(ctx context.Context, req *authv1.CreateRequest) (*authv1.CreateResponse, error) {
	user := converter.ProtoAddUser2AddUser(req)

	id, err := i.authService.AddUser(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("error when saving the user: %w", err)
	}
	return &authv1.CreateResponse{Id: *id}, nil
}
