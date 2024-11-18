package config

import (
	"errors"
	"os"
	"strconv"
	"time"
)

// JWTConfig структура реализующая конфиг для jwt.
type JWTConfig struct {
	refreshTokenSecretKey  string
	accessTokenSecretKey   string
	accessTokenExpiration  string
	refreshTokenExpiration string
}

// JwtConfig интерфейс для получения элементов из конфига.
type JwtConfig interface {
	GetRefreshTokenSecretKey() string
	GetAccessTokenSecretKey() string
	GetAccessTokenExpiration() (time.Duration, error)
	GetRefreshTokenExpiration() (time.Duration, error)
}

// NewJWTConfig конструктор для конфига jwt.
func NewJWTConfig() (JwtConfig, error) {
	RefreshTokenSecretKey := os.Getenv("REFRESH_TOKEN_SECRET_KEY")
	if RefreshTokenSecretKey == "" {
		return nil, errors.New("REFRESH_TOKEN_SECRET_KEY environment variable not set")
	}

	AccessTokenSecretKey := os.Getenv("ACCESS_TOKEN_SECRET_KEY")
	if AccessTokenSecretKey == "" {
		return nil, errors.New("ACCESS_TOKEN_SECRET_KEY environment variable not set")
	}

	AccessTokenExpiration := os.Getenv("ACCESS_TOKEN_EXPIRATION")
	if AccessTokenExpiration == "" {
		return nil, errors.New("ACCESS_TOKEN_EXPIRATION environment variable not set")
	}

	RefreshTokenExpiration := os.Getenv("REFRESH_TOKEN_EXPIRATION")
	if RefreshTokenExpiration == "" {
		return nil, errors.New("REFRESH_TOKEN_EXPIRATION environment variable not set")
	}

	return &JWTConfig{
		RefreshTokenSecretKey,
		AccessTokenSecretKey,
		AccessTokenExpiration,
		RefreshTokenExpiration,
	}, nil
}

// GetRefreshTokenSecretKey получаем ключ для подписи refresh.
func (j JWTConfig) GetRefreshTokenSecretKey() string {
	return j.refreshTokenSecretKey
}

// GetAccessTokenSecretKey получаем ключ для подписи access.
func (j JWTConfig) GetAccessTokenSecretKey() string {
	return j.accessTokenSecretKey
}

// GetAccessTokenExpiration получаем время жизни токена access.
func (j JWTConfig) GetAccessTokenExpiration() (time.Duration, error) {
	CountMinute, err := strconv.Atoi(j.accessTokenExpiration)
	if err != nil {
		return 0, err
	}
	return time.Duration(CountMinute) * time.Minute, nil
}

// GetRefreshTokenExpiration получаем время жизни токена refresh.
func (j JWTConfig) GetRefreshTokenExpiration() (time.Duration, error) {
	CountMinute, err := strconv.Atoi(j.refreshTokenExpiration)
	if err != nil {
		return 0, err
	}
	return time.Duration(CountMinute) * time.Minute, nil
}
