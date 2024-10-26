package auth

import (
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
	authv1 "github.com/Dnlbb/auth/pkg/auth_v1"
)

type Implementation struct {
	authv1.UnimplementedAuthServer
	authService servinterfaces.AuthService
}

func NewImplementation(authService servinterfaces.AuthService) *Implementation {
	return &Implementation{authService: authService}
}
