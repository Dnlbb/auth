package dao

import (
	auth "github.com/Dnlbb/auth/pkg/auth_v1"
	"time"
)

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
