package auth

import (
	"context"
	"fmt"

	"github.com/Dnlbb/auth/internal/converter"
	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
)

func (i *Implementation) Get(ctx context.Context, req *authv1.GetRequest) (*authv1.GetResponse, error) {
	params, err := converter.GetUserParamsReq2Params(req)
	if err != nil {
		return nil, fmt.Errorf("get request params: %v", err)
	}
	userProfile, err := i.authService.GetUser(ctx, *params)
	if err != nil {
		return nil, fmt.Errorf("get user: %v", err)
	}
	response := converter.UserModel2ProtoUserProfile(*userProfile)
	return response, nil
}
