package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
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
		ID        int64     `db:"id" redis:"id" json:"id"`
		Name      string    `db:"name" json:"name"`
		Email     string    `db:"email" json:"email"`
		Password  string    `db:"password" json:"password"`
		Role      string    `db:"role" json:"role"`
		CreatedAt time.Time `db:"created_at" json:"created_at"`
		UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
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

	// UserPayload полезная нагрузка в jwt токен.
	UserPayload struct {
		Username string `json:"username"`
		Role     string `json:"role"`
	}

	// UserClaims claims в jwt токене.
	UserClaims struct {
		jwt.StandardClaims
		Username string `json:"username"`
		Role     string `json:"role"`
	}
)
