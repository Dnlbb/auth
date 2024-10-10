package dao

import (
	"github.com/Dnlbb/auth/pkg/auth"
)

// PostgresInterface Далее будет заменен для гибкой работы с бд
type PostgresInterface interface {
	Save(user User) error
	Update(update UpdateUser) error
	Delete(id DeleteID) error
	Get(id GetID) (User, error)
}
type (
	// User for db
	User struct {
		Name     string
		Email    string
		Role     auth.Role
		Password string
	}
	// UpdateUser for db
	UpdateUser struct {
		ID    int64
		Name  string
		Email string
		Role  auth.Role
	}
	// DeleteID for db
	DeleteID int64
	// GetID for db
	GetID int64
)
