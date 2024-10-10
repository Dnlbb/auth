package dao

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
	"log"
	"os"
)

const (
	dbDSN = "DEFAULT_DSN"
)

type Storage struct {
	con *pgx.Conn
	ctx context.Context
}

func InitStorage() (*Storage, error) {
	ctx := context.Background()
	dbDSN := GetDSN()
	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	return &Storage{con: con, ctx: ctx}, nil
}

func GetDSN() string {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = dbDSN
	}
	return dsn
}

func (s *Storage) CloseCon() error {
	err := s.con.Close(s.ctx)
	return err
}

func (s *Storage) Save(user User) error {
	res, err := s.con.Exec(s.ctx, "INSERT INTO USERS (name, email, role, password) VALUES ($1, $2, $3, $4)", user.Name, user.Email, user.Role, user.Password)
	if err != nil {
		return errors.Wrap(err, "Error inserting user")
	}
	log.Printf("Inserted user: %v", res.RowsAffected())
	return nil
}

func (s *Storage) Update(update UpdateUser) error {
	res, err := s.con.Exec(s.ctx, "UPDATE USERS SET name = $1, email = $2, role = $3 WHERE id = $4", update.Name, update.Email, update.Role, update.Id)
	if err != nil {
		return errors.Wrap(err, "Error updating user")
	}
	log.Printf("Updated user: %v", res.RowsAffected())
	return nil
}

func (s *Storage) Delete(id DeleteId) error {
	res, err := s.con.Exec(s.ctx, "DELETE FROM USERS WHERE id = $1", id)
	if err != nil {
		return errors.Wrap(err, "Error deleting user")
	}
	log.Printf("Deleted user: %v", res.RowsAffected())
	return nil
}

func (s *Storage) Get(id GetId) (User, error) {
	res, err := s.con.Query(s.ctx, "SELECT name, email, role FROM USERS WHERE id = $1", id)
	if err != nil {
		errors.Wrap(err, "Error getting user")
	}
	defer res.Close()
	var resUser User
	err = res.Scan(&resUser.Name, &resUser.Email, &resUser.Role)
	if err != nil {
		errors.Wrap(err, "Error getting user")
	}
	log.Printf("Got user: %+v", resUser)
	return resUser, err
}
