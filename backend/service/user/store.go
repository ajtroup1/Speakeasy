package user

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/ajtroup1/speakeasy/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found with id '%d'", id)
	}

	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	firstname := capitalizeFirstLetter(user.Firstname)
	lastname := capitalizeFirstLetter(user.Lastname)

	_, err := s.db.Exec("INSERT INTO users (username, password, firstname, lastname, email, phoneNumber, imglink, status, createdAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.Username, user.Password, firstname, lastname, user.Email, user.PhoneNumber, user.ImgLink, user.Status, user.CreatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) EditUser(user types.User) error {
	firstname := capitalizeFirstLetter(user.Firstname)
	lastname := capitalizeFirstLetter(user.Lastname)

	_, err := s.db.Exec("UPDATE users SET username = ?, firstname = ?, lastname = ?, email = ?, phoneNumber = ?, imgLink = ? WHERE id = ?",
		user.Username, firstname, lastname, user.Email, user.PhoneNumber, user.ImgLink, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Firstname, &user.Lastname, &user.Email, &user.PhoneNumber, &user.ImgLink, &user.Status, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}
