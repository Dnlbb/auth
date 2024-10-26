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

	authImpl *auth.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

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

func (s *serviceProvider) GetDBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.GetPGConfig().DSN())
		if err != nil {
			log.Fatal("failed to connect to database: %w", err)
		}
		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatal("failed to ping database: %w", err)
		}
		closer.Add(cl.Close)
		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) GetTxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.GetDBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) GetAuthRepository(ctx context.Context) repoInterface.StorageInterface {
	if s.authRepository == nil {
		s.authRepository = storage.NewPostgresRepo(s.GetDBClient(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) GetAuthService(ctx context.Context) servinterfaces.AuthService {
	if s.authService == nil {
		s.authService = authserv.NewService(s.GetAuthRepository(ctx),
			s.GetTxManager(ctx),
		)
	}

	return s.authService
}

func (s *serviceProvider) GetAuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.GetAuthService(ctx))
	}

	return s.authImpl
}
