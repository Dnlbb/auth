package user

import (
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
	userv1 "github.com/Dnlbb/auth/pkg/user_v1"
)

// Controller структура реализующая сгенерированный grpc сервер
type Controller struct {
	userv1.UnimplementedUserApiServer
	authService servinterfaces.AuthService
}

// NewController конструктор для реализации grpc сервера
func NewController(authService servinterfaces.AuthService) *Controller {
	return &Controller{authService: authService}
}
