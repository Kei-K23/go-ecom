package users

import (
	"database/sql"
	"fmt"

	"github.com/Kei-K23/go-ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}
func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	// Prepare SQL statement
	stmt, err := s.db.Prepare("SELECT * FROM users WHERE email = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query and scan the result into a User struct
	var u types.User
	err = stmt.QueryRow(email).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, err
	}

	return &u, nil
}
func (s *Store) GetUserByID(id int) (*types.User, error) {
	// Prepare SQL statement
	stmt, err := s.db.Prepare("SELECT * FROM users WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Execute query and scan the result into a User struct
	var u types.User
	err = stmt.QueryRow(id).Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Password, &u.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	} else if err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *Store) CreateUser(user types.User) error {
	stmt, err := s.db.Prepare("INSERT INTO users (first_name, last_name, email, password) VALUES (?, ?, ?, ?)")

	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}
