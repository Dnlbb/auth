package dao

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
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
	res, err := s.con.Exec(s.ctx, "INSERT INTO USERS (name, email, role, password) VALUES ($1, $2, $3, $4)", user.Name, user.Email, user.Role, user.Password)
	if err != nil {
		log.Fatal("Error inserting user into database")
		return err
	}
	log.Printf("Inserted user: %v", res.RowsAffected())
	return nil
}

// Update for postgresql
func (s *Storage) Update(update UpdateUser) error {
	res, err := s.con.Exec(s.ctx, "UPDATE USERS SET name = $1, email = $2, role = $3 WHERE id = $4", update.Name, update.Email, update.Role, update.ID)
	if err != nil {
		log.Println("Error updating user")
		return err
	}
	log.Printf("Updated user: %v", res.RowsAffected())
	return nil
}

// Delete for postgresql
func (s *Storage) Delete(id DeleteID) error {
	res, err := s.con.Exec(s.ctx, "DELETE FROM USERS WHERE id = $1", id)
	if err != nil {
		log.Fatal("Error deleting user")
		return err
	}
	log.Printf("Deleted user: %v", res.RowsAffected())
	return nil
}

// Get for postgresql
func (s *Storage) Get(id GetID) (User, error) {
	res, err := s.con.Query(s.ctx, "SELECT name, email, role FROM USERS WHERE id = $1", id)
	if err != nil {
		log.Fatal(err, "Error getting user")
	}
	defer res.Close()
	var resUser User
	err = res.Scan(&resUser.Name, &resUser.Email, &resUser.Role)
	if err != nil {
		log.Fatal(err, "Error getting user")
	}
	log.Printf("Got user: %+v", resUser)
	return resUser, err
}
