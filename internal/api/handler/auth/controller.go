package auth

import (
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
	"github.com/Dnlbb/auth/pkg/auth_v1"
)

// Controller контроллер для api авторизации.
type Controller struct {
	auth_v1.UnimplementedAuthServer
	authorizationService servinterfaces.AuthorizationService
}

// NewControllerAuthorization конструктор для контроллера api авторизации.
func NewControllerAuthorization(authorizationService servinterfaces.AuthorizationService) *Controller {
	return &Controller{
		authorizationService: authorizationService,
	}
}
