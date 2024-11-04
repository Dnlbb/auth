package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/Dnlbb/auth/internal/api/auth"
	"github.com/Dnlbb/auth/internal/models"
	serviceMocks "github.com/Dnlbb/auth/internal/service/mocks"
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
	"github.com/Dnlbb/auth/pkg/auth_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	type (
		authServiceMockFunc func(mc *minimock.Controller) servinterfaces.AuthService
		args                struct {
			ctx context.Context
			req *auth_v1.CreateRequest
		}
	)

	var (
		id       = gofakeit.Int64()
		ctx      = context.Background()
		mc       = minimock.NewController(t)
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, true, 1)

		errorCreate = errors.New("error with service")
		res         = &auth_v1.CreateResponse{
			Id: id,
		}
	)

	defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            *auth_v1.CreateResponse
		err             error
		authServiceMock authServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: &auth_v1.CreateRequest{
					User: &auth_v1.User{
						Name:  name,
						Email: email,
						Role:  auth_v1.Role_USER,
					},
					Password:        password,
					PasswordConfirm: password,
				},
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) servinterfaces.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				// Не вынесено отдельно, чтобы проверить конвертацию разных ролей.
				mock.CreateMock.Expect(ctx, models.User{
					Name:     name,
					Email:    email,
					Password: password,
					Role:     "USER",
				}).Return(&id, nil)
				return mock
			},
		},
		{
			name: "success case: check convert role admin",
			args: args{
				ctx: ctx,
				req: &auth_v1.CreateRequest{
					User: &auth_v1.User{
						Name:  name,
						Email: email,
						Role:  auth_v1.Role_ADMIN,
					},
					Password:        password,
					PasswordConfirm: password,
				},
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) servinterfaces.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				// Не вынесено отдельно, чтобы проверить конвертацию разных ролей.
				mock.CreateMock.Expect(ctx, models.User{
					Name:     name,
					Email:    email,
					Password: password,
					Role:     "ADMIN",
				}).Return(&id, nil)
				return mock
			},
		},
		{
			name: "success case: check convert role unspecified",
			args: args{
				ctx: ctx,
				req: &auth_v1.CreateRequest{
					User: &auth_v1.User{
						Name:  name,
						Email: email,
						Role:  auth_v1.Role_ROLE_UNSPECIFIED,
					},
					Password:        password,
					PasswordConfirm: password,
				},
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) servinterfaces.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				// Не вынесено отдельно, чтобы проверить конвертацию разных ролей.
				mock.CreateMock.Expect(ctx, models.User{
					Name:     name,
					Email:    email,
					Password: password,
					Role:     "ROLE_UNSPECIFIED",
				}).Return(&id, nil)
				return mock
			},
		},
		{
			name: "error case: error in the service layer",
			args: args{
				ctx: ctx,
				req: &auth_v1.CreateRequest{
					User: &auth_v1.User{
						Name:  name,
						Email: email,
						Role:  auth_v1.Role_USER,
					},
					Password:        password,
					PasswordConfirm: password,
				},
			},
			want: nil,
			err:  fmt.Errorf("error while creating: %w", errorCreate),
			authServiceMock: func(mc *minimock.Controller) servinterfaces.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.CreateMock.Expect(ctx, models.User{
					Name:     name,
					Email:    email,
					Password: password,
					Role:     "USER",
				}).Return(nil, errorCreate)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			AuthServiceMock := tt.authServiceMock(mc)
			api := auth.NewController(AuthServiceMock)

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}

			require.Equal(t, tt.want, newID)
		})
	}
}
