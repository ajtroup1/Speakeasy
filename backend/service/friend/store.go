package friend

import (
	"database/sql"
	"fmt"

	"github.com/ajtroup1/speakeasy/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetFriendshipByIDs(sendID, receiveID uint) (bool, error) {
	query := "SELECT COUNT(*) FROM friends WHERE sendingID = ? AND receivingID = ?"

	var count int
	err := s.db.QueryRow(query, sendID, receiveID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to get friendship: %w", err)
	}

	// Return true if the friendship exists, otherwise false
	return count > 0, nil
}

func (s *Store) GetFriendshipsByID(userID uint) ([]*types.User, error) {
	query := `
		SELECT u.id, u.username, u.password, u.firstname, u.lastname, u.email, u.phoneNumber, u.imgLink, u.status, u.createdAt, u.textNotifications, u.emailNotifications
		FROM friends f
		JOIN users u ON u.id = f.receivingID
		WHERE f.sendingID = ? AND f.status = 2
		UNION
		SELECT u.id, u.username, u.password, u.firstname, u.lastname, u.email, u.phoneNumber, u.imgLink, u.status, u.createdAt, u.textNotifications, u.emailNotifications
		FROM friends f
		JOIN users u ON u.id = f.sendingID
		WHERE f.receivingID = ? AND f.status = 2
	`

	rows, err := s.db.Query(query, userID, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get friendships: %w", err)
	}
	defer rows.Close()

	var friends []*types.User
	for rows.Next() {
		user, err := scanRowsIntoUser(rows)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		friends = append(friends, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	if len(friends) < 1 {
		return nil, fmt.Errorf("no (accepted) friends found for user id: %d", userID)
	}

	return friends, nil
}

func (s *Store) FriendUser(sendID, receiveID uint) error {
	_, err := s.db.Exec("INSERT INTO friends (sendingID, receivingID, status) VALUES (?, ?, ?)", sendID, receiveID, 1)
	if err != nil {
		return fmt.Errorf("failed to add friend: %w", err)
	}
	return nil
}

func (s *Store) Accept(sendID, receiveID uint) error {
	_, err := s.db.Exec("UPDATE friends SET status = ? WHERE sendingID = ? AND receivingID = ?", 2, sendID, receiveID)
	if err != nil {
		return fmt.Errorf("failed to accept friend request: %w", err)
	}
	return nil
}

func (s *Store) UnfriendUser(sendID, receiveID uint) error {
	_, err := s.db.Exec("UPDATE friends SET status = ? WHERE sendingID = ? AND receivingID = ?", 0, sendID, receiveID)
	if err != nil {
		return fmt.Errorf("failed to unfriend user: %w", err)
	}
	return nil
}

func (s *Store) Refriend(sendID, receiveID uint) error {
	_, err := s.db.Exec("UPDATE friends SET status = ? WHERE sendingID = ? AND receivingID = ?", 1, sendID, receiveID)
	if err != nil {
		return fmt.Errorf("failed to friend user: %w", err)
	}
	return nil
}

func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Firstname, &user.Lastname, &user.Email, &user.PhoneNumber, &user.ImgLink, &user.Status, &user.CreatedAt, &user.TextNotifications, &user.EmailNotifications)
	if err != nil {
		return nil, err
	}

	return user, nil
}
