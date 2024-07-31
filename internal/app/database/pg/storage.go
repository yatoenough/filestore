package pg

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db              *sqlx.DB
	imageRepository *ImageRepository
}

func New(connStr string) *Storage {
	db := sqlx.MustConnect("pgx", connStr)
	return &Storage{db: db}
}

func (s *Storage) Close() error {
	err := s.db.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Image() *ImageRepository {
	if s.imageRepository != nil {
		return s.imageRepository
	}

	s.imageRepository = &ImageRepository{
		storage: s,
	}

	return s.imageRepository
}
