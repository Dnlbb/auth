package authserv

import (
	"github.com/Dnlbb/auth/internal/producer"
	"github.com/Dnlbb/auth/internal/repository/repointerface"
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
	"github.com/Dnlbb/platform_common/pkg/db"
)

type service struct {
	cache         repointerface.CacheInterface
	storage       repointerface.StorageInterface
	txManager     db.TxManager
	kafkaProducer producer.Producer
}

// NewService конструктор сервиса
func NewService(storage repointerface.StorageInterface,
	txManager db.TxManager,
	cache repointerface.CacheInterface,
	kafkaProducer producer.Producer) servinterfaces.AuthService {
	return &service{storage: storage,
		txManager:     txManager,
		cache:         cache,
		kafkaProducer: kafkaProducer,
	}
}
