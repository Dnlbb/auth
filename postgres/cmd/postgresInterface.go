package dao

// PostgresInterface Далее будет заменен для гибкой работы с бд
type PostgresInterface interface {
	Save(user User) error
	Update(update UpdateUser) error
	Delete(id DeleteID) error
	GetUser(params GetUserParams) (*User, error)
}
