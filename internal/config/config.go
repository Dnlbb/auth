package config

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

// LoadEnv загружаем переменные окружения из файла auth.env в окружение проекта.
func LoadEnv(path2env string) error {
	if err := godotenv.Load(path2env); err != nil {
		return fmt.Errorf("error loading auth.env file: %w, path to env: %s", err, path2env)
	}

	return nil
}

// GRPCConfig интерфейс получения адреса запуска сервера.
type GRPCConfig interface {
	Address() string
}

// PGConfig интерфейс получения DSN для старта хранилища.
type PGConfig interface {
	DSN() string
}

// RedisConfig интерфейс для получения данных для конфига Redis.
type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
}
