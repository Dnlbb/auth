package tests

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	clientMocks "github.com/Dnlbb/auth/internal/client/mocks"
	"github.com/Dnlbb/auth/internal/models"
	repoMocks "github.com/Dnlbb/auth/internal/repository/mocks"
	"github.com/Dnlbb/auth/internal/repository/repointerface"
	"github.com/Dnlbb/auth/internal/service/authserv"
	"github.com/Dnlbb/platform_common/pkg/db"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type (
		AuthCacheMockFunc   func(mc *minimock.Controller) repointerface.CacheInterface
		AuthTxManMockFunc   func(mc *minimock.Controller) db.TxManager
		AuthStorageMockFunc func(mc *minimock.Controller) repointerface.StorageInterface
	)

	type args struct {
		ctx context.Context
		req models.User
	}

	var (
		ctx            = context.Background()
		mc             = minimock.NewController(t)
		name           = gofakeit.Name()
		correctEmail   = "Dr.Pepper@gmail.com"
		password       = "12345678910"
		id             = gofakeit.Int64()
		errPass        = errors.New("invalid password: password must be at least 8 characters but no more 255")
		errTransaction = errors.New("error transaction")
		errCache       = errors.New("error cache")
	)

	defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            *int64
		err             error
		authCacheMock   AuthCacheMockFunc
		authTxManMock   AuthTxManMockFunc
		authStorageMock AuthStorageMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: models.User{
					Name:     name,
					Email:    correctEmail,
					Password: password,
					Role:     "USER",
				},
			},
			want: &id,
			err:  nil,
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				mock.CreateMock.Return(nil)
				return mock
			},
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Return(nil)
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				return mock
			},
		},
		{
			name: "error case: long password",
			args: args{
				ctx: ctx,
				req: models.User{
					Name:     name,
					Email:    correctEmail,
					Password: strings.Repeat("a", 300),
					Role:     "USER",
				},
			},
			want: nil,
			err:  errPass,
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				return mock
			},
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				return mock
			},
		},
		{
			name: "error case: short password",
			args: args{
				ctx: ctx,
				req: models.User{
					Name:     name,
					Email:    correctEmail,
					Password: "abc",
					Role:     "USER",
				},
			},
			want: nil,
			err:  errPass,
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				return mock
			},
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				return mock
			},
		},
		{
			name: "error case: transaction error",
			args: args{
				ctx: ctx,
				req: models.User{
					Name:     name,
					Email:    correctEmail,
					Password: password,
					Role:     "USER",
				},
			},
			want: nil,
			err:  fmt.Errorf("user creation error: %w", errTransaction),
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				return mock
			},
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Return(errTransaction)
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				return mock
			},
		},
		{
			name: "error case: err with cache",
			args: args{
				ctx: ctx,
				req: models.User{
					Name:     name,
					Email:    correctEmail,
					Password: password,
					Role:     "USER",
				},
			},
			want: &id,
			err:  nil,
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				mock.CreateMock.Return(errCache)
				return mock
			},
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Return(nil)
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			RepoMock := tt.authStorageMock(mc)
			CacheMock := tt.authCacheMock(mc)
			TxManMock := tt.authTxManMock(mc)
			service := authserv.NewService(RepoMock, TxManMock, CacheMock)

			_, err := service.Create(tt.args.ctx, tt.args.req)
			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
