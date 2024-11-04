package redis

import (
	"github.com/Dnlbb/auth/internal/client/cache/redis"
	"github.com/Dnlbb/auth/internal/repository/repointerface"
)

type cache struct {
	cl redis.Client
}

// NewRedisCache конструктор для redis.
func NewRedisCache(cl redis.Client) repointerface.CacheInterface {
	return &cache{cl: cl}
}
