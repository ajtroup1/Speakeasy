package channel

import (
	"database/sql"

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

func (s *Store) CreateChannel(channel types.Channel) error {
	_, err := s.db.Exec("INSERT INTO channels (name, description, createdAt, createdBy, imgLink) VALUES (?, ?, ?, ?, ?)",
		channel.Name, channel.Description, channel.CreatedAt, channel.CreatedBy, channel.ImgLink)
	return err
}