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
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDelete(t *testing.T) {
	type authServiceMockFunc func(mc *minimock.Controller) servinterfaces.AuthService

	type args struct {
		ctx context.Context
		req *auth_v1.DeleteRequest
	}

	var (
		idToDelete  = gofakeit.Int64()
		ctx         = context.Background()
		mc          = minimock.NewController(t)
		errorDelete = errors.New("error with delete service")
		res         = &emptypb.Empty{}
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
			name: "success case",
			args: args{
				ctx: ctx,
				req: &auth_v1.DeleteRequest{
					Id: idToDelete,
				},
			},
			want: res,
			err:  nil,
			authServiceMock: func(mc *minimock.Controller) servinterfaces.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.DeleteMock.Expect(ctx, models.DeleteID(idToDelete)).Return(nil)
				return mock
			},
		},
		{
			name: "error case: error in the service layer",
			args: args{
				ctx: ctx,
				req: &auth_v1.DeleteRequest{
					Id: idToDelete,
				},
			},
			want: &emptypb.Empty{},
			err:  fmt.Errorf("deleting user error: %w", errorDelete),
			authServiceMock: func(mc *minimock.Controller) servinterfaces.AuthService {
				mock := serviceMocks.NewAuthServiceMock(mc)
				mock.DeleteMock.Expect(ctx, models.DeleteID(idToDelete)).Return(errorDelete)
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

			resp, err := api.Delete(tt.args.ctx, tt.args.req)
			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.want, resp)
		})
	}
}
