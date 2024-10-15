package dao

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

const (
	dbDSN = "DEFAULT_DSN"
)

// Storage storage
type Storage struct {
	con *pgx.Conn
	ctx context.Context
}

// InitStorage storage for postgresql
func InitStorage() (*Storage, error) {
	ctx := context.Background()
	dbDSN := getDSN()
	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	return &Storage{con: con, ctx: ctx}, nil
}

func getDSN() string {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = dbDSN
	}
	return dsn
}

// CloseCon close connection to db
func (s *Storage) CloseCon() {
	err := s.con.Close(s.ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// Save for postgresql
func (s *Storage) Save(user User) error {
	res, err := s.con.Exec(s.ctx, "INSERT INTO users (name, email, role, password) VALUES ($1, $2, $3, $4)", user.Name, user.Email, user.Role, user.Password)
	if err != nil {
		return fmt.Errorf("error inserting user into database: %w", err)
	}
	log.Printf("Inserted user: %v", res.RowsAffected())
	return nil
}

// Update for postgresql
func (s *Storage) Update(update UpdateUser) error {
	res, err := s.con.Exec(s.ctx, "UPDATE users SET name = $1, email = $2, role = $3, password = $4 WHERE id = $5", update.Name, update.Email, update.Role, update.Password, update.ID)
	if err != nil {
		return fmt.Errorf("error updating user into database: %w", err)
	}
	log.Printf("Updated user: %v", res.RowsAffected())
	return nil
}

// Delete for postgresql
func (s *Storage) Delete(id DeleteID) error {
	res, err := s.con.Exec(s.ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting user into database: %w", err)
	}
	log.Printf("Deleted user: %v", res.RowsAffected())
	return nil
}

// GetUser for postgresql
func (s *Storage) GetUser(params GetUserParams) (*User, error) {
	var user User
	var err error
	query := sq.Select("id", "name", "email", "role", "created_at", "updated_at").From("users")
	switch {
	case params.ID != nil:
		query = query.Where(sq.Eq{"id": *params.ID})
	case params.Username != nil:
		query = query.Where(sq.Eq{"username": *params.Username})
	default:
		return nil, fmt.Errorf("не указан ни ID, ни Username")
	}
	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error building sql query: %w", err)
	}
	row := s.con.QueryRow(s.ctx, sqlQuery, args...)
	err = row.Scan(&user.ID,
		&user.Name,
		&user.Email,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, fmt.Errorf("пользователя не существует")
	} else if err != nil {
		return nil, fmt.Errorf("ошибка при обращении в базу для получения профиля пользователя: %v", err)
	}
	return &user, nil
}
