package models

import (
	"errors"
)

var (
	// ErrUserNotFound нет пользователя в хранилище.
	ErrUserNotFound = errors.New("user not found")
)
