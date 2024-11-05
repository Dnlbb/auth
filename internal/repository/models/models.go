package models

type (
	// User модель пользователя для redis.
	User struct {
		ID        string `redis:"id"`
		Name      string `redis:"name"`
		Email     string `redis:"email"`
		Password  string `redis:"password"`
		Role      string `redis:"role"`
		CreatedAt string `redis:"created_at"`
		UpdatedAt string `redis:"updated_at"`
	}
)
