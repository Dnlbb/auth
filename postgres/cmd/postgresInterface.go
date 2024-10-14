package dao

import (
	"time"

	auth "github.com/Dnlbb/auth/pkg/auth_v1"
)

// PostgresInterface Далее будет заменен для гибкой работы с бд
type PostgresInterface interface {
	Save(user User) error
	Update(update UpdateUser) error
	Delete(id DeleteID) error
	GetUser(params GetUserParams) (*User, error)
}
type (
	// User for db
	User struct {
		ID        int64
		Name      string
		Email     string
		Password  string
		Role      auth.Role
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	// UpdateUser for db
	UpdateUser struct {
		ID       int64
		Name     string
		Email    string
		Password string
		Role     auth.Role
	}
	// DeleteID for db
	DeleteID int64
	// GetUserParams for db
	GetUserParams struct {
		ID       *int64
		Username *string
	}
)
