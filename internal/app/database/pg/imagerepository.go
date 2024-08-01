package pg

import (
	"github.com/yatoenough/filestore/internal/app/model"
)

type ImageRepository struct {
	storage *Storage
}

func (r *ImageRepository) Create(i *model.Image) (*model.Image, error) {
	res, err := r.storage.db.NamedQuery(
		"INSERT INTO images (original_name, store_name, path, ext) VALUES (:original_name, :store_name, :path, :ext) RETURNING id",
		&i,
	)

	if err != nil {
		return nil, err
	}

	if res.Next() {
		res.Scan(&i.ID)
	}

	return i, nil
}

func (r *ImageRepository) FindAll() (*[]model.Image, error) {
	var images []model.Image
	err := r.storage.db.Select(&images, "SELECT id, original_name, store_name, path, ext FROM images")
	if err != nil {
		return nil, err
	}

	return &images, nil
}

func (r *ImageRepository) FindById(id int64) (*model.Image, error) {
	var image model.Image
	err := r.storage.db.Get(&image, "SELECT id, original_name, store_name, path, ext FROM images WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &image, nil
}
