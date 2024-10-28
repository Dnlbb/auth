package auth

import (
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
)

// Controller структура реализующая сгенерированный grpc сервер
type Controller struct {
	authv1.UnimplementedAuthServer
	authService servinterfaces.AuthService
}

// NewController конструктор для реализации grpc сервера
func NewController(authService servinterfaces.AuthService) *Controller {
	return &Controller{authService: authService}
}
