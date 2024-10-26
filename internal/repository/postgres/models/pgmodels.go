package pgmodels

import (
	"time"
)

// Log_key тип для логирования запросов в базу
type LogKey string

// Ключи для логирования запросов в базу
const (
	SAVE    LogKey = "save"
	UPDATE  LogKey = "update"
	DELETE  LogKey = "delete"
	GETUSER LogKey = "getuser"
)

type (
	// UserAdd модель для регистрации пользователя
	UserAdd struct {
		Name     string
		Email    string
		Password string
		Role     string
	}

	// User моделька для сохранения пользователя в базу данных, а также для получения информации о пользователе из базы данных.
	User struct {
		ID        int64     `db:"id"`
		Name      string    `db:"name"`
		Email     string    `db:"email"`
		Password  string    `db:"password"`
		Role      string    `db:"role"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
	// UpdateUser моделька для обновления данных пользователя, (name, email, password, role) меняются каждый раз в базе при передаче этих значений.
	UpdateUser struct {
		ID       int64  `db:"id"`
		Name     string `db:"name"`
		Email    string `db:"email"`
		Password string `db:"password"`
		Role     string `db:"role"`
	}
	// UserID модель для получения UserID
	UserID struct {
		ID int64 `db:"id"`
	}
	// DeleteID моделька для удаления пользователя по его id из базы данных.
	DeleteID int64
	// GetUserParams моделька для получения информации о пользователе, в ручке предусмотрена логика вариативной передачи аргументов,
	// то есть мы передаем либо id для получения профиля, либо имя пользователя.
	GetUserParams struct {
		ID       *int64
		Username *string
	}
)
