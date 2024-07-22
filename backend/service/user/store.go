package user

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/ajtroup1/speakeasy/service/auth"
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

func (s *Store) GetAllUsers() ([]*types.User, error) {
	rows, err := s.db.Query("SELECT id, username, password, firstname, lastname, email, phoneNumber, imgLink, status, createdAt, textNotifications, emailNotifications FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*types.User
	for rows.Next() {
		var u types.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.Firstname, &u.Lastname, &u.Email, &u.PhoneNumber, &u.ImgLink, &u.Status, &u.CreatedAt, &u.TextNotifications, &u.EmailNotifications); err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("users not found")
	}

	return users, nil
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

	_, err := s.db.Exec("INSERT INTO users (username, password, firstname, lastname, email, phoneNumber, imglink, status, createdAt, textNotifications, emailNotifications) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.Username, user.Password, firstname, lastname, user.Email, user.PhoneNumber, user.ImgLink, user.Status, user.CreatedAt, user.TextNotifications, user.EmailNotifications)
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

func (s *Store) ChangePassword(id uint, currentPassword, newPassword, confirmNewPassword string) error {
	if newPassword != confirmNewPassword {
		return fmt.Errorf("new password did not match confirm new password")
	}

	// Retrieve the user
	row := s.db.QueryRow("SELECT password FROM users WHERE id = ?", id)
	var storedPassword string
	err := row.Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("user not found with id '%d'", id)
		}
		return err
	}

	// Check the current password
	if !auth.ComparePasswords(storedPassword, []byte(currentPassword)) {
		return fmt.Errorf("current password did not match")
	}

	// Check if new password is the same as the current one
	if auth.ComparePasswords(storedPassword, []byte(newPassword)) {
		return fmt.Errorf("new password cannot be the same as current password")
	}

	// Hash the new password
	hashedPassword, err := auth.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("error hashing new password: %v", err)
	}

	// Update the password in the database
	_, err = s.db.Exec("UPDATE users SET password = ? WHERE id = ?", hashedPassword, id)
	if err != nil {
		return err
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

	err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Firstname, &user.Lastname, &user.Email, &user.PhoneNumber, &user.ImgLink, &user.Status, &user.CreatedAt, &user.TextNotifications, &user.EmailNotifications)
	if err != nil {
		return nil, err
	}

	return user, nil
}
