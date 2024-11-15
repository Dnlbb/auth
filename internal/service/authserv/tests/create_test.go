package tests

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	clientMocks "github.com/Dnlbb/auth/internal/client/mocks"
	"github.com/Dnlbb/auth/internal/models"
	"github.com/Dnlbb/auth/internal/producer"
	"github.com/Dnlbb/auth/internal/producer/mocks"
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
		AuthCacheMockFunc    func(mc *minimock.Controller) repointerface.CacheInterface
		AuthTxManMockFunc    func(mc *minimock.Controller) db.TxManager
		AuthStorageMockFunc  func(mc *minimock.Controller) repointerface.StorageInterface
		AuthProducerMockFunc func(mc *minimock.Controller) producer.Producer
	)

	type args struct {
		ctx context.Context
		req models.User
	}

	var (
		part         = int32(1)
		offset       = int64(3)
		ctx          = context.Background()
		mc           = minimock.NewController(t)
		name         = gofakeit.Name()
		correctEmail = "Dr.Pepper@gmail.com"
		password     = "12345678910"
		id           = gofakeit.Int64()
		errPass      = errors.New("invalid password: password must be at least 8 characters but no more 255")
		errSave      = errors.New("error save")
		errLog       = errors.New("error log")
		errCache     = errors.New("error cache")
		user         = models.User{
			Name:     name,
			Email:    correctEmail,
			Password: password,
			Role:     "USER",
		}
	)

	defer mc.Finish()
	tests := []struct {
		name             string
		args             args
		want             *int64
		err              error
		authCacheMock    AuthCacheMockFunc
		authTxManMock    AuthTxManMockFunc
		authStorageMock  AuthStorageMockFunc
		authProducerMock AuthProducerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: user,
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
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				mock.SaveMock.Expect(ctx, user).Return(id, nil)
				mock.LogMock.Expect(ctx, models.SAVE).Return(nil)
				return mock
			},
			authProducerMock: func(mc *minimock.Controller) producer.Producer {
				mock := mocks.NewProducerMock(mc)
				mock.SendMessageMock.Return(part, offset, nil)
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
			authProducerMock: func(mc *minimock.Controller) producer.Producer {
				mock := mocks.NewProducerMock(mc)
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
			authProducerMock: func(mc *minimock.Controller) producer.Producer {
				mock := mocks.NewProducerMock(mc)
				return mock
			},
		},
		{
			name: "error case: save user error",
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
			err:  fmt.Errorf("user creation error: %w", errSave),
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
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
				mock.SaveMock.Expect(ctx, user).Return(0, errSave)
				return mock
			},
			authProducerMock: func(mc *minimock.Controller) producer.Producer {
				mock := mocks.NewProducerMock(mc)
				return mock
			},
		},
		{
			name: "error case: err with log",
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
			err:  fmt.Errorf("user creation error: %w", errLog),
			authCacheMock: func(mc *minimock.Controller) repointerface.CacheInterface {
				mock := repoMocks.NewCacheInterfaceMock(mc)
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
				mock.SaveMock.Expect(ctx, user).Return(id, nil)
				mock.LogMock.Expect(ctx, models.SAVE).Return(errLog)
				return mock
			},
			authProducerMock: func(mc *minimock.Controller) producer.Producer {
				mock := mocks.NewProducerMock(mc)
				return mock
			},
		},
		{
			name: "success case: err with cache",
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
				mock.ReadCommittedMock.Set(func(ctx context.Context, handler db.Handler) error {
					return handler(ctx)
				})
				return mock
			},
			authStorageMock: func(mc *minimock.Controller) repointerface.StorageInterface {
				mock := repoMocks.NewStorageInterfaceMock(mc)
				mock.SaveMock.Expect(ctx, user).Return(id, nil)
				mock.LogMock.Expect(ctx, models.SAVE).Return(nil)
				return mock
			},
			authProducerMock: func(mc *minimock.Controller) producer.Producer {
				mock := mocks.NewProducerMock(mc)
				mock.SendMessageMock.Return(part, offset, nil)
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
			Producer := tt.authProducerMock(mc)
			service := authserv.NewService(RepoMock, TxManMock, CacheMock, Producer)

			_, err := service.Create(tt.args.ctx, tt.args.req)
			if tt.err != nil {
				require.EqualError(t, err, tt.err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
