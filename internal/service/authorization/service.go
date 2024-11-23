package authorization

import (
	"github.com/Dnlbb/auth/internal/config"
	"github.com/Dnlbb/auth/internal/repository/repointerface"
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
)

type service struct {
	cache        repointerface.CacheInterface
	storage      repointerface.StorageInterface
	accessPolicy repointerface.AccessPolicies
	config       config.JwtConfig
}

// NewService конструктор сервиса
func NewService(storage repointerface.StorageInterface,
	cache repointerface.CacheInterface,
	accessPolicy repointerface.AccessPolicies,
	config config.JwtConfig) servinterfaces.AuthorizationService {
	return &service{
		storage:      storage,
		cache:        cache,
		accessPolicy: accessPolicy,
		config:       config,
	}
}
