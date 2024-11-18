package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/Dnlbb/auth/internal/api/user"
	"github.com/Dnlbb/auth/internal/models"
	serviceMocks "github.com/Dnlbb/auth/internal/service/mocks"
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
	"github.com/Dnlbb/auth/pkg/user_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdate(t *testing.T) {
	type authServiceMockFunc func(mc *minimock.Controller) servinterfaces.AuthService

	type args struct {
		ctx context.Context
		req *user_v1.UpdateRequest
	}

	var (
		idToUpdate  = gofakeit.Int64()
		ctx         = context.Background()
		mc          = minimock.NewController(t)
		name        = gofakeit.Name()
		email       = gofakeit.Email()
		errorUpdate = errors.New("error updating user")
	)

	defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case: update role to USER",
			args: args{
				ctx: ctx,
				req: &user_v1.UpdateRequest{
					Id:    idToUpdate,
					Name:  wrapperspb.String(name),
					Email: wrapperspb.String(email),
					Role:  user_v1.Role_USER,
				},
			},
			want: &emptypb.Empty{},
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) servinterfaces.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.UpdateMock.Expect(ctx, models.User{
					ID:    idToUpdate,
					Name:  name,
					Email: email,
					Role:  "USER",
				}).Return(nil)
				return mock
			},
		},
		{
			name: "success case: convert role ROLE_UNSPECIFIED",
			args: args{
				ctx: ctx,
				req: &user_v1.UpdateRequest{
					Id:    idToUpdate,
					Name:  wrapperspb.String(name),
					Email: wrapperspb.String(email),
					Role:  user_v1.Role_ROLE_UNSPECIFIED,
				},
			},
			want: &emptypb.Empty{},
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) servinterfaces.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.UpdateMock.Expect(ctx, models.User{
					ID:    idToUpdate,
					Name:  name,
					Email: email,
					Role:  "ROLE_UNSPECIFIED",
				}).Return(nil)
				return mock
			},
		},
		{
			name: "success case: update role to ADMIN",
			args: args{
				ctx: ctx,
				req: &user_v1.UpdateRequest{
					Id:    idToUpdate,
					Name:  wrapperspb.String(name),
					Email: wrapperspb.String(email),
					Role:  user_v1.Role_ADMIN,
				},
			},
			want: &emptypb.Empty{},
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) servinterfaces.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.UpdateMock.Expect(ctx, models.User{
					ID:    idToUpdate,
					Name:  name,
					Email: email,
					Role:  "ADMIN",
				}).Return(nil)
				return mock
			},
		},
		{
			name: "error case: service layer error",
			args: args{
				ctx: ctx,
				req: &user_v1.UpdateRequest{
					Id:    idToUpdate,
					Name:  wrapperspb.String(name),
					Email: wrapperspb.String(email),
					Role:  user_v1.Role_USER,
				},
			},
			want: &emptypb.Empty{},
			err:  fmt.Errorf("error updating user: %w", errorUpdate),
			authServiceMock: func(mc *minimock.Controller) servinterfaces.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.UpdateMock.Expect(ctx, models.User{
					ID:    idToUpdate,
					Name:  name,
					Email: email,
					Role:  "USER",
				}).Return(errorUpdate)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			AuthServiceMock := tt.authServiceMock(mc)
			api := user.NewController(AuthServiceMock)

			result, err := api.Update(tt.args.ctx, tt.args.req)
			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want, result)
		})
	}
}
