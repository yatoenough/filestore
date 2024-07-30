package pg

import (
	"log"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sqlx.DB
}

func New(connStr string) (*Storage, error) {
	db, err := sqlx.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	log.Println("Disconnecting database...")
	err := s.db.Close()
	if err != nil {
		return err
	}
	log.Println("Disconnected")
	return nil
}
