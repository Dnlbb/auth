package authserv

import (
	"github.com/Dnlbb/auth/internal/client/db"
	"github.com/Dnlbb/auth/internal/repository/repoInterface"
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
)

type service struct {
	storage   repoInterface.StorageInterface
	txManager db.TxManager
}

func NewService(storage repoInterface.StorageInterface,
	txManager db.TxManager) servinterfaces.AuthService {
	return &service{storage: storage,
		txManager: txManager}
}
