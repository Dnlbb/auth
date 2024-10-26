package auth

import (
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
)

// Implementation структура реализующая сгенерированный grpc сервер
type Implementation struct {
	authv1.UnimplementedAuthServer
	authService servinterfaces.AuthService
}

// NewImplementation конструктор для реализации grpc сервера
func NewImplementation(authService servinterfaces.AuthService) *Implementation {
	return &Implementation{authService: authService}
}
