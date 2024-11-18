package access

import (
	"github.com/Dnlbb/auth/internal/service/servinterfaces"
	"github.com/Dnlbb/auth/pkg/access_v1"
)

// Controller контроллер для api проверки доступа.
type Controller struct {
	access_v1.UnimplementedAccessServer
	accessService servinterfaces.AccessService
}

// NewController конструктор для контроллера для api проверки доступа.
func NewController(accessService servinterfaces.AccessService) *Controller {
	return &Controller{
		accessService: accessService,
	}
}
