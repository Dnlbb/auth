package tests_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	clientMocks "github.com/Dnlbb/auth/internal/client/mocks"
	"github.com/Dnlbb/auth/internal/models"
	repoMocks "github.com/Dnlbb/auth/internal/repository/mocks"
	"github.com/Dnlbb/auth/internal/repository/repointerface"
	"github.com/Dnlbb/auth/internal/service/userserv"
	"github.com/Dnlbb/platform_common/pkg/db"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type (
		AuthCacheMockFunc   func(mc *minimock.Controller) repointerface.CacheInterface
		AuthTxManMockFunc   func(mc *minimock.Controller) db.TxManager
		AuthStorageMockFunc func(mc *minimock.Controller) repointerface.StorageInterface
		args                struct {
			ctx    context.Context
			params models.GetUserParams
		}
	)

	var (
		ctx             = context.Background()
		mc              = minimock.NewController(t)
		errCacheMiss    = models.ErrUserNotFound
		errCacheGeneric = errors.New("cache error")
		errLogFailure   = errors.New("log error")
		errStorage      = errors.New("storage error")
		errCaching      = errors.New("caching error")
		email           = gofakeit.Email()
		password        = gofakeit.PetName()
		userID          = int64(123)
		username        = "testuser"
		userProfile     = &models.User{ID: userID,
			Name:      username,
			Email:     email,
			Password:  password,
			Role:      "USER",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	)

	defer t.Cleanup(mc.Finish)
	tests := []struct {
		name            string
		args            args
		want            *models.User
		err             error
		authCacheMock   AuthCacheMockFunc
		authTxManMock   AuthTxManMockFunc
		authStorageMock AuthStorageMockFunc
	}{
		{
			name: "success case: user found in cache",
			args: args{
				ctx:    ctx,
				params: models.GetUserParams{ID: &userID},
			},
			want: userProfile,
			err:  nil,
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				mock.GetMock.Return(userProfile, nil)
				return mock
			},
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				return repoMocks.NewStorageInterfaceMock(mc)
			},
		},
		{
			name: "error case: cache error",
			args: args{
				ctx:    ctx,
				params: models.GetUserParams{ID: &userID},
			},
			want: nil,
			err:  fmt.Errorf("error with cache: %w", errCacheGeneric),
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				mock.GetMock.Return(nil, errCacheGeneric)
				return mock
			},
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				return repoMocks.NewStorageInterfaceMock(mc)
			},
		},
		{
			name: "error case (cache miss): transaction error storage",
			args: args{
				ctx:    ctx,
				params: models.GetUserParams{ID: &userID},
			},
			want: nil,
			err:  fmt.Errorf("error getting user profile: %w", errStorage),
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				mock.GetMock.Return(nil, errCacheMiss)
				return mock
			},
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				mock.GetUserMock.Expect(ctx, models.GetUserParams{ID: &userID}).Return(nil, errStorage)
				return mock
			},
		},
		{
			name: "error case (cache miss): transaction error logging",
			args: args{
				ctx:    ctx,
				params: models.GetUserParams{ID: &userID},
			},
			want: nil,
			err:  fmt.Errorf("error logging: %w", errLogFailure),
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				mock.GetMock.Return(nil, errCacheMiss)
				return mock
			},
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				mock.GetUserMock.Expect(ctx, models.GetUserParams{ID: &userID}).Return(userProfile, nil)
				mock.LogMock.Expect(ctx, models.GETUSER).Return(errLogFailure)
				return mock
			},
		},
		{
			name: "error case (cache miss): transaction error caching",
			args: args{
				ctx:    ctx,
				params: models.GetUserParams{ID: &userID},
			},
			want: nil,
			err:  fmt.Errorf("error caching user profile: %w", errCaching),
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
				mock.GetMock.Return(nil, errCacheMiss)
				mock.CreateMock.Expect(ctx, userID, *userProfile).Return(errCaching)
				return mock
			},
			authTxManMock: func(mc *minimock.Controller) db.TxManager {
				mock := clientMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				mock.GetUserMock.Expect(ctx, models.GetUserParams{ID: &userID}).Return(userProfile, nil)
				mock.LogMock.Expect(ctx, models.GETUSER).Return(nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			CacheMock := tt.authCacheMock(mc)
			TxManMock := tt.authTxManMock(mc)
			StorageMock := tt.authStorageMock(mc)
			service := userserv.NewService(StorageMock, TxManMock, CacheMock, nil)

			result, err := service.Get(tt.args.ctx, tt.args.params)
			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
				require.Nil(t, result)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, result)
			}
		})
	}
}
