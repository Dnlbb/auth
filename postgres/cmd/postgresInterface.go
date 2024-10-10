package dao

import (
	"github.com/Dnlbb/auth/pkg/auth"
)

type PostgresInterface interface {
	Save(user User) error
	Update(update UpdateUser) error
	Delete(id DeleteId) error
	Get(id GetId) (User, error)
}
type (
	User struct {
		Name     string
		Email    string
		Role     auth.Role
		Password string
	}

	UpdateUser struct {
		Id    int64
		Name  string
		Email string
		Role  auth.Role
	}

	DeleteId int64

	GetId int64
)
