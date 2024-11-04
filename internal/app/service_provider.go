package app

import (
	"context"
	"log"

	"github.com/Dnlbb/auth/internal/api/auth"
	"github.com/Dnlbb/auth/internal/client/cache/redis"
	"github.com/Dnlbb/auth/internal/config"
	"github.com/Dnlbb/auth/internal/repository/postgres/storage"
	redisCache "github.com/Dnlbb/auth/internal/repository/redis"
	"github.com/Dnlbb/auth/internal/repository/repointerface"
	"github.com/Dnlbb/auth/internal/service/authserv"
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
	"github.com/Dnlbb/platform_common/pkg/closer"
	"github.com/Dnlbb/platform_common/pkg/db"
	"github.com/Dnlbb/platform_common/pkg/db/pg"
	"github.com/Dnlbb/platform_common/pkg/db/transaction"
	redigo "github.com/gomodule/redigo/redis"
)

type serviceProvider struct {
	pgConfig    config.PGConfig
	grpcConfig  config.GRPCConfig
	redisConfig config.RedisConfig

	dbClient    db.Client
	redisPool   *redigo.Pool
	redisClient redis.Client
	txManager   db.TxManager

	serviceCache   repointerface.CacheInterface
	authRepository repointerface.StorageInterface

	authService servinterfaces.AuthService

	authController *auth.Controller
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// GetPGConfig получаем конфиг постгреса, для последующего получения DSN.
func (s *serviceProvider) GetPGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPgConfig()
		if err != nil {
			log.Fatal("failed to load pg config: %w", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

// GetGRPCConfig получаем конфиг grpc, для последующего получения адреса нашего приложения и старта на нем.
func (s *serviceProvider) GetGRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGrpcConfig()
		if err != nil {
			log.Fatal("failed to load gRPC config: %w", err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// GetRedisConfig получаем конфиг для redis.
func (s *serviceProvider) GetRedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := config.NewRedisConfig()
		if err != nil {
			log.Fatal("failed to load redis config: %w", err)
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

// GetDBClient инициализируем клиента к базе данных.
func (s *serviceProvider) GetDBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.GetPGConfig().DSN())
		if err != nil {
			log.Fatal("failed to connect to database: %w", err)
		}

		if err = cl.DB().Ping(ctx); err != nil {
			log.Fatal("failed to ping database: %w", err)
		}

		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) RedisPool() *redigo.Pool {
	if s.redisPool == nil {
		s.redisPool = &redigo.Pool{
			MaxIdle:     s.GetRedisConfig().MaxIdle(),
			IdleTimeout: s.GetRedisConfig().IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", s.GetRedisConfig().Address())
			},
		}
	}

	return s.redisPool
}

func (s *serviceProvider) GetRedisClient(_ context.Context) redis.Client {
	if s.redisConfig == nil {
		s.redisClient = redis.NewClient(s.RedisPool(), s.GetRedisConfig())
	}

	return s.redisClient
}

// GetTxManager инициализация менеджера транзакций.
func (s *serviceProvider) GetTxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.GetDBClient(ctx).DB())
	}

	return s.txManager
}

// GetAuthRepository инициализация хранилища.
func (s *serviceProvider) GetAuthRepository(ctx context.Context) repointerface.StorageInterface {
	if s.authRepository == nil {
		s.authRepository = storage.NewPostgresRepo(s.GetDBClient(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) GetCache(ctx context.Context) repointerface.CacheInterface {
	if s.serviceCache == nil {
		s.serviceCache = redisCache.NewRedisCache(s.GetRedisClient(ctx))
	}

	return s.serviceCache
}

// GetAuthService инициализация сервиса авторизации.
func (s *serviceProvider) GetAuthService(ctx context.Context) servinterfaces.AuthService {
	if s.authService == nil {
		s.authService = authserv.NewService(s.GetAuthRepository(ctx),
			s.GetTxManager(ctx),
			s.GetCache(ctx),
		)
	}

	return s.authService
}

// GetAuthController инициализация контроллера.
func (s *serviceProvider) GetAuthController(ctx context.Context) *auth.Controller {
	if s.authController == nil {
		s.authController = auth.NewController(s.GetAuthService(ctx))
	}

	return s.authController
}
