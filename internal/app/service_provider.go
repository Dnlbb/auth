package app

import (
	"context"
	"log"

	"github.com/Dnlbb/auth/internal/api/auth"
	"github.com/Dnlbb/auth/internal/client/db"
	"github.com/Dnlbb/auth/internal/client/db/pg"
	"github.com/Dnlbb/auth/internal/client/db/transaction"
	"github.com/Dnlbb/auth/internal/closer"
	"github.com/Dnlbb/auth/internal/config"
	"github.com/Dnlbb/auth/internal/repository/postgres/storage"
	"github.com/Dnlbb/auth/internal/repository/repoInterface"
	"github.com/Dnlbb/auth/internal/service/authserv"
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	dbClient       db.Client
	txManager      db.TxManager
	authRepository repoInterface.StorageInterface

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

// GetTxManager инициализация менеджера транзакций.
func (s *serviceProvider) GetTxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.GetDBClient(ctx).DB())
	}

	return s.txManager
}

// GetAuthRepository инициализация хранилища.
func (s *serviceProvider) GetAuthRepository(ctx context.Context) repoInterface.StorageInterface {
	if s.authRepository == nil {
		s.authRepository = storage.NewPostgresRepo(s.GetDBClient(ctx))
	}

	return s.authRepository
}

// GetAuthService инициализация сервиса авторизации.
func (s *serviceProvider) GetAuthService(ctx context.Context) servinterfaces.AuthService {
	if s.authService == nil {
		s.authService = authserv.NewService(s.GetAuthRepository(ctx),
			s.GetTxManager(ctx),
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
