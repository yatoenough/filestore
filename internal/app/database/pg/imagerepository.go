package pg

import (
	"github.com/yatoenough/filestore/internal/app/model"
)

type ImageRepository struct {
	storage *Storage
}

func (r *ImageRepository) Create(i *model.Image) (*model.Image, error) {
	if err := r.storage.db.QueryRow(
		"INSERT INTO images (original_name, store_name, path, ext) VALUES ($1, $2, $3, $4) RETURNING id",
		i.OriginalName,
		i.StoreName,
		i.Path,
		i.Extension,
	).Scan(&i.ID); err != nil {
		return nil, err
	}

	return i, nil
}
