package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/yatoenough/filestore/internal/app/model"
)

type ImageRepository struct {
	storage *sqlx.DB
}

func NewImageRepo(s *sqlx.DB) *ImageRepository {
	return &ImageRepository{
		storage: s,
	}
}

func (r *ImageRepository) Create(i *model.Image) (*model.Image, error) {
	res, err := r.storage.NamedQuery(
		"INSERT INTO images (original_name, store_name, path, ext) VALUES (:original_name, :store_name, :path, :ext) RETURNING id",
		&i,
	)

	if err != nil {
		return nil, err
	}

	res.Scan(&i.ID)

	return i, nil
}

func (r *ImageRepository) FindAll() (*[]model.Image, error) {
	var images []model.Image
	err := r.storage.Select(&images, "SELECT id, original_name, store_name, path, ext FROM images")
	if err != nil {
		return nil, err
	}

	return &images, nil
}

func (r *ImageRepository) FindById(id int64) (*model.Image, error) {
	var image model.Image
	err := r.storage.Get(&image, "SELECT id, original_name, store_name, path, ext FROM images WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &image, nil
}

func (r *ImageRepository) Delete(id int64) error {
	_, err := r.storage.Exec("DELETE FROM images WHERE id = $1", id)
	if err != nil {
		return err
	}

	return nil
}
