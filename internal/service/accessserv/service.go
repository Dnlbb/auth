package accessserv

import (
	"github.com/Dnlbb/auth/internal/config"
	"github.com/Dnlbb/auth/internal/repository/repointerface"
)

// Service сервис для проверки доступа.
type Service struct {
	accessPolicy repointerface.AccessPolicies
	config       config.JwtConfig
}

// NewService конструктор для сервиса.
func NewService(accessPolicy repointerface.AccessPolicies) *Service {
	return &Service{
		accessPolicy: accessPolicy,
	}
}
