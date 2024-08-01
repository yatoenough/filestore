package pg

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/yatoenough/filestore/internal/app/database/repository"
)

type Storage struct {
	db              *sqlx.DB
	imageRepository *repository.ImageRepository
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

func (s *Storage) Image() *repository.ImageRepository {
	if s.imageRepository != nil {
		return s.imageRepository
	}

	s.imageRepository = repository.NewImageRepo(s.db)

	return s.imageRepository
}
