package app

import (
	"context"
	"log"

	"github.com/Dnlbb/auth/internal/api/handler/auth"
	"github.com/Dnlbb/auth/internal/api/handler/user"
	"github.com/Dnlbb/auth/internal/client/cache/redis"
	"github.com/Dnlbb/auth/internal/config"
	"github.com/Dnlbb/auth/internal/repository/AccessPolicies"
	"github.com/Dnlbb/auth/internal/repository/postgres/storage"
	redisCache "github.com/Dnlbb/auth/internal/repository/redis"
	"github.com/Dnlbb/auth/internal/repository/repointerface"
	userService "github.com/Dnlbb/auth/internal/service/user"

	"github.com/Dnlbb/auth/internal/service/authorization"
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
	"github.com/Dnlbb/platform_common/pkg/closer"
	"github.com/Dnlbb/platform_common/pkg/db"
	"github.com/Dnlbb/platform_common/pkg/db/pg"
	"github.com/Dnlbb/platform_common/pkg/db/transaction"
	"github.com/IBM/sarama"
	redigo "github.com/gomodule/redigo/redis"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	redisConfig   config.RedisConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConf
	kafkaConfig   config.KafkaConf
	jwtConfig     config.JwtConfig

	kafkaProducer sarama.SyncProducer
	dbClient      db.Client
	redisPool     *redigo.Pool
	redisClient   redis.Client
	txManager     db.TxManager

	serviceCache   repointerface.CacheInterface
	userRepository repointerface.StorageInterface
	accessPolicy   repointerface.AccessPolicies

	userService          servinterfaces.UserService
	authorizationService servinterfaces.AuthorizationService

	userController          *user.Controller
	authorizationController *auth.Controller
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

// GetJwtConfig получаем конфиг для jwt.
func (s *serviceProvider) GetJwtConfig() config.JwtConfig {
	if s.jwtConfig == nil {
		cfg, err := config.NewJWTConfig()
		if err != nil {
			log.Fatal("failed to load JWT config: %w", err)
		}

		s.jwtConfig = cfg
	}

	return s.jwtConfig
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

func (s *serviceProvider) GetHTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatal("failed to load http config: %w", err)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) GetSwaggerConfig() config.SwaggerConf {
	if s.swaggerConfig == nil {
		cfg, err := config.NewSwaggerServerConf()
		if err != nil {
			log.Fatal("failed to load swagger config: %w", err)
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
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

func (s *serviceProvider) GetkafkaConfig() config.KafkaConf {
	if s.kafkaConfig == nil {
		cfg, err := config.NewKafkaConfImpl()
		if err != nil {
			log.Fatal("failed to load kafka config: %w", err)
		}

		s.kafkaConfig = cfg
	}

	return s.kafkaConfig
}

func (s *serviceProvider) GetKafkaProducer() sarama.SyncProducer {
	if s.kafkaProducer == nil {
		conf := sarama.NewConfig()
		conf.Producer.RequiredAcks = sarama.WaitForAll
		conf.Producer.Return.Successes = true
		conf.Producer.Retry.Max = 5

		s.GetkafkaConfig()

		producer, err := sarama.NewSyncProducer(s.kafkaConfig.Addresses(), conf)
		if err != nil {
			log.Fatal("failed to create kafka producer: %w", err)
		}

		closer.Add(producer.Close)

		s.kafkaProducer = producer
	}

	return s.kafkaProducer
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

// GetUserRepository инициализация хранилища.
func (s *serviceProvider) GetUserRepository(ctx context.Context) repointerface.StorageInterface {
	if s.userRepository == nil {
		s.userRepository = storage.NewPostgresRepo(s.GetDBClient(ctx))
	}

	return s.userRepository
}

// GetAccessPolicyRepository инициализация политик доступа.
func (s *serviceProvider) GetAccessPolicyRepository(_ context.Context) repointerface.AccessPolicies {
	if s.accessPolicy == nil {
		s.accessPolicy = AccessPolicies.NewAccessPolicyRepository()
	}

	return s.accessPolicy
}

func (s *serviceProvider) GetCache(ctx context.Context) repointerface.CacheInterface {
	if s.serviceCache == nil {
		s.serviceCache = redisCache.NewRedisCache(s.GetRedisClient(ctx))
	}

	return s.serviceCache
}

// GetUserService инициализация сервиса авторизации.
func (s *serviceProvider) GetUserService(ctx context.Context) servinterfaces.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.GetUserRepository(ctx),
			s.GetTxManager(ctx),
			s.GetCache(ctx),
			s.GetKafkaProducer(),
		)
	}

	return s.userService
}

func (s *serviceProvider) GetAuthorizationService(ctx context.Context) servinterfaces.AuthorizationService {
	if s.authorizationService == nil {
		s.authorizationService = authorization.NewService(
			s.GetUserRepository(ctx),
			s.GetCache(ctx),
			s.GetAccessPolicyRepository(ctx),
			s.GetJwtConfig())
	}

	return s.authorizationService
}

// GetUserController инициализация контроллера.
func (s *serviceProvider) GetUserController(ctx context.Context) *user.Controller {
	if s.userController == nil {
		s.userController = user.NewController(s.GetUserService(ctx))
	}

	return s.userController
}

func (s *serviceProvider) GetAuthorizationController(ctx context.Context) *auth.Controller {
	if s.authorizationController == nil {
		s.authorizationController = auth.NewControllerAuthorization(s.GetAuthorizationService(ctx))
	}

	return s.authorizationController
}
