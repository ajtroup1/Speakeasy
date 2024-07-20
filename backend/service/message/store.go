package message

import (
	"database/sql"
	"fmt"

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

func (s *Store) GetMessageByID(id int) (*types.Message, error) {
	rows, err := s.db.Query("SELECT * FROM messages WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, fmt.Errorf("message not found with id '%d'", id)
	}

	m, err := scanRowsIntoMessage(rows)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *Store) CreateMessage(message types.Message) error {
	_, err := s.db.Exec("INSERT INTO messages (content, createdAt, createdBy, createdIn) VALUES (?, ?, ?, ?)",
		message.Content, message.CreatedAt, message.CreatedBy, message.ChannelD)
	return err
}

func scanRowsIntoMessage(rows *sql.Rows) (*types.Message, error) {
	message := new(types.Message)

	err := rows.Scan(&message.ID, &message.Content, &message.CreatedAt, &message.CreatedBy, &message.ChannelD)
	if err != nil {
		return nil, err
	}

	return message, nil
}
