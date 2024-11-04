package models

import (
	"time"
)

// LogKey тип для логирования запросов в базу.
type LogKey string

// Ключи для логирования запросов в базу.
const (
	SAVE    LogKey = "save"
	UPDATE  LogKey = "update"
	DELETE  LogKey = "delete"
	GETUSER LogKey = "getuser"
)

type (
	// User моделька профиля пользователя для сервисного слоя.
	User struct {
		ID        int64     `db:"id" redis:"id"`
		Name      string    `db:"name" redis:"name"`
		Email     string    `db:"email" redis:"email"`
		Password  string    `db:"password" redis:"password"`
		Role      string    `db:"role" redis:"role"`
		CreatedAt time.Time `db:"created_at" redis:"created_at"`
		UpdatedAt time.Time `db:"updated_at" redis:"updated_at"`
	}
	// UserID id пользователя.
	UserID int64

	// DeleteID моделька для удаления пользователя по его, сервисный слой.
	DeleteID int64

	// GetUserParams моделька для получения информации о пользователе, в ручке предусмотрена логика вариативной передачи аргументов,
	// то есть мы передаем либо id для получения профиля, либо имя пользователя, сервисный слой.
	GetUserParams struct {
		ID       *int64
		Username *string
	}
)
