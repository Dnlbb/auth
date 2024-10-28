package storage

import (
	"github.com/Dnlbb/auth/internal/client/db"
	"github.com/Dnlbb/auth/internal/repository/repoInterface"
)

type storage struct {
	db db.Client
}

// NewPostgresRepo инициализируем хранилище postgresql и приводим его к типу интерфейса StorageInterface.
func NewPostgresRepo(db db.Client) repoInterface.StorageInterface {
	return &storage{db: db}
}
