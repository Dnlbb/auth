package tests

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/Dnlbb/auth/internal/api/auth"
	"github.com/Dnlbb/auth/internal/models"
	serviceMocks "github.com/Dnlbb/auth/internal/service/mocks"
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
	"github.com/Dnlbb/auth/pkg/auth_v1"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {
	type (
		authServiceMockFunc func(mc *minimock.Controller) servinterfaces.AuthService
		args                struct {
			ctx context.Context
			req *auth_v1.GetRequest
		}
	)

	var (
		ctx       = context.Background()
		mc        = minimock.NewController(t)
		errorGet  = errors.New("service get error")
		id        = int64(123)
		username  = "test"
		name      = "Ivan"
		email     = "Ivan@example.com"
		createdAt = time.Now()
		updatedAt = time.Now()
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *auth_v1.GetResponse
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case: get by id",
			args: args{
				ctx: ctx,
				req: &auth_v1.GetRequest{
					NameOrId: &auth_v1.GetRequest_Id{Id: id},
				},
			},
			want: &auth_v1.GetResponse{
				Id: id,
				User: &auth_v1.User{
					Name:  name,
					Email: email,
					Role:  auth_v1.Role_USER,
				},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			},
			err: nil,
			authServiceMock: func(mc *minimock.Controller) servinterfaces.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetMock.Expect(ctx, models.GetUserParams{ID: &id}).Return(&models.User{
					ID:        id,
					Name:      name,
					Email:     email,
					Role:      "USER",
					CreatedAt: createdAt,
					UpdatedAt: updatedAt,
				}, nil)
				return mock
			},
		},
		{
			name: "success case: get by username",
			args: args{
				ctx: ctx,
				req: &auth_v1.GetRequest{
					NameOrId: &auth_v1.GetRequest_Username{Username: username},
				},
			},
			want: &auth_v1.GetResponse{
				Id: id,
				User: &auth_v1.User{
					Name:  name,
					Email: email,
					Role:  auth_v1.Role_USER,
				},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt),
			},
			err: nil,
			authServiceMock: func(mc *minimock.Controller) servinterfaces.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetMock.Expect(ctx, models.GetUserParams{Username: &username}).Return(&models.User{
					ID:        id,
					Name:      name,
					Email:     email,
					Role:      "USER",
					CreatedAt: createdAt,
					UpdatedAt: updatedAt,
				}, nil)
				return mock
			},
		},
		{
			name: "error case: service get error",
			args: args{
				ctx: ctx,
				req: &auth_v1.GetRequest{
					NameOrId: &auth_v1.GetRequest_Id{Id: id},
				},
			},
			want: nil,
			err:  fmt.Errorf("error when getUser request: %w", errorGet),
			authServiceMock: func(mc *minimock.Controller) servinterfaces.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.GetMock.Expect(ctx, models.GetUserParams{ID: &id}).Return(nil, errorGet)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			authServiceMock := tt.authServiceMock(mc)
			api := auth.NewController(authServiceMock)

			resp, err := api.Get(tt.args.ctx, tt.args.req)
			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want, resp)
		})
	}
}
