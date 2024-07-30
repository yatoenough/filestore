package pg

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func New(connStr string) (*Storage, error) {
	db := sqlx.MustConnect("pgx", connStr)
	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	err := s.db.Close()
	if err != nil {
		return err
	}
	return nil
}
