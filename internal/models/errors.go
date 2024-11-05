package models

import (
	"errors"
)

var (
	// ErrUserNotFound нет его в кэше.
	ErrUserNotFound = errors.New("user not found")
)
