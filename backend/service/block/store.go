package block

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

func (s *Store) GetBlockByIDs(sendID, receiveID uint) (bool, error) {
    query := "SELECT COUNT(*) FROM blocks WHERE sendingID = ? AND receivingID = ?"

    var count int
    err := s.db.QueryRow(query, sendID, receiveID).Scan(&count)
    if err != nil {
        return false, fmt.Errorf("failed to get blocks: %w", err)
    }

    // Return true if the block exists, otherwise false
    return count > 0, nil
}

func (s *Store) BlockUser(sendID, receiveID uint) error {
    _, err := s.db.Exec("INSERT INTO blocks (sendingID, receivingID, status) VALUES (?, ?, ?)", sendID, receiveID, 1)
    if err != nil {
        return fmt.Errorf("failed to block user: %w", err)
    }
    return nil
}

func (s *Store) UnblockUser(sendID, receiveID uint) error {
    _, err := s.db.Exec("UPDATE blocks SET status = ? WHERE sendingID = ? AND receivingID = ?", 0, sendID, receiveID)
    if err != nil {
        return fmt.Errorf("failed to unblock user: %w", err)
    }
    return nil
}

func (s *Store) Reblock(sendID, receiveID uint) error {
    _, err := s.db.Exec("UPDATE blocks SET status = ? WHERE sendingID = ? AND receivingID = ?", 1, sendID, receiveID)
    if err != nil {
        return fmt.Errorf("failed to block user: %w", err)
    }
    return nil
}