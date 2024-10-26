package models

import (
	"time"
)

type (
	// UserAdd модель для регистрации пользователя
	UserAdd struct {
		Name     string
		Email    string
		Password string
		Role     string
	}

	// User моделька профиля пользователя для сервисного слоя.
	User struct {
		ID        int64
		Name      string
		Email     string
		Password  string
		Role      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	// UpdateUser моделька для обновления данных пользователя, в сервисном слое.
	UpdateUser struct {
		ID       int64
		Name     string
		Email    string
		Password string
		Role     string
	}
	// DeleteID моделька для удаления пользователя по его, сервисный слой.
	DeleteID int64
	// GetUserParams моделька для получения информации о пользователе, в ручке предусмотрена логика вариативной передачи аргументов,
	// то есть мы передаем либо id для получения профиля, либо имя пользователя, сервисный слой.
	GetUserParams struct {
		ID       *int64
		Username *string
	}
)
