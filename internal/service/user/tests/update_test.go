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
	"github.com/Dnlbb/auth/internal/service/user"
	"github.com/Dnlbb/platform_common/pkg/db"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestUpdate(t *testing.T) {
	t.Parallel()

	type (
		AuthTxManMockFunc   func(mc *minimock.Controller) db.TxManager
		AuthStorageMockFunc func(mc *minimock.Controller) repointerface.StorageInterface
		AuthCacheMockFunc   func(mc *minimock.Controller) repointerface.CacheInterface
	)

	var (
		ctx        = context.Background()
		mc         = minimock.NewController(t)
		userUpdate = models.User{ID: 1,
			Name:     "testuser",
			Email:    "test@test.com",
			Password: "12345678910"}
		errUpdate = errors.New("update error")
		errLog    = errors.New("log error")
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		userUpdate      models.User
		err             error
		authTxManMock   AuthTxManMockFunc
		authStorageMock AuthStorageMockFunc
		authCacheMock   AuthCacheMockFunc
	}{
		{
			name:       "success case",
			userUpdate: userUpdate,
			err:        nil,
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				mock.UpdateMock.Expect(ctx, userUpdate).Return(nil)
				mock.LogMock.Expect(ctx, models.UPDATE).Return(nil)
				return mock
			},
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				return mock
			},
		},
		{
			name:       "error case: update error",
			userUpdate: userUpdate,
			err:        fmt.Errorf("could not update user: %w", errUpdate),
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				mock.UpdateMock.Expect(ctx, userUpdate).Return(errUpdate)
				return mock
			},
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				return mock
			},
		},
		{
			name:       "error case: log error",
			userUpdate: userUpdate,
			err:        fmt.Errorf("could not update user: %w", errLog),
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				mock.UpdateMock.Expect(ctx, userUpdate).Return(nil)
				mock.LogMock.Expect(ctx, models.UPDATE).Return(errLog)
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
			service := user.NewService(RepoMock, TxManMock, CacheMock, nil)

			err := service.Update(ctx, tt.userUpdate)
			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
