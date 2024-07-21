package friend

import (
	"database/sql"
	"fmt"
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

func (s *Store) FriendUser(sendID, receiveID uint) error {
    _, err := s.db.Exec("INSERT INTO friends (sendingID, receivingID, status) VALUES (?, ?, ?)", sendID, receiveID, 1)
    if err != nil {
        return fmt.Errorf("failed to add friend: %w", err)
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