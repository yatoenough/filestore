package service

import (
	"github.com/yatoenough/filestore/internal/app/database/repository"
	"github.com/yatoenough/filestore/internal/app/model"
)

type Repo interface {
	Create(i *model.Image) (*model.Image, error)
}

type ImageService struct {
	repo Repo
}

func NewImageService(r *repository.ImageRepository) *ImageService {
	return &ImageService{
		repo: r,
	}
}
