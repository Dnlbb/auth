package tests

import (
	"context"
	"errors"
	"fmt"
	"testing"

	clientMocks "github.com/Dnlbb/auth/internal/client/mocks"
	"github.com/Dnlbb/auth/internal/models"
	repoMocks "github.com/Dnlbb/auth/internal/repository/mocks"
	"github.com/Dnlbb/auth/internal/repository/repointerface"
	"github.com/Dnlbb/auth/internal/service/authserv"
	"github.com/Dnlbb/platform_common/pkg/db"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestDelete(t *testing.T) {
	t.Parallel()

	type (
		AuthTxManMockFunc   func(mc *minimock.Controller) db.TxManager
		AuthStorageMockFunc func(mc *minimock.Controller) repointerface.StorageInterface
		AuthCacheMockFunc   func(mc *minimock.Controller) repointerface.CacheInterface
	)

	var (
		ctx       = context.Background()
		mc        = minimock.NewController(t)
		userID    = models.DeleteID(1)
		errDelete = errors.New("delete error")
		errLog    = errors.New("log error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		userID          models.DeleteID
		err             error
		authTxManMock   AuthTxManMockFunc
		authStorageMock AuthStorageMockFunc
		authCacheMock   AuthCacheMockFunc
	}{
		{
			name:   "success case",
			userID: userID,
			err:    nil,
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				mock.DeleteMock.Expect(ctx, userID).Return(nil)
				mock.LogMock.Expect(ctx, models.DELETE).Return(nil)
				return mock
			},
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				return mock
			},
		},
		{
			name:   "error case: delete error",
			userID: userID,
			err:    fmt.Errorf("error deleting user: %w", errDelete),
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				mock.DeleteMock.Expect(ctx, userID).Return(errDelete)
				return mock
			},
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				return mock
			},
		},
		{
			name:   "error case: log error",
			userID: userID,
			err:    fmt.Errorf("error deleting user: %w", errLog),
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				mock.DeleteMock.Expect(ctx, userID).Return(nil)
				mock.LogMock.Expect(ctx, models.DELETE).Return(errLog)
				return mock
			},
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			RepoMock := tt.authStorageMock(mc)
			TxManMock := tt.authTxManMock(mc)
			CacheMock := tt.authCacheMock(mc)
			service := authserv.NewService(RepoMock, TxManMock, CacheMock, nil)

			err := service.Delete(ctx, tt.userID)
			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
