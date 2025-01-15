package service

import (
	"apw/internal/domain/models"
	repo "apw/internal/repository"
)

type UploadService struct {
	repo repo.Upload
}

func NewUploadService(repo repo.Upload) *UploadService {
	return &UploadService{repo: repo}
}

func (s *UploadService) Create(upload models.Upload) (int, error) {
	return s.repo.Create(upload)
}

func (s *UploadService) GetAll() ([]models.Upload, error) {
	return s.repo.GetAll()
}

func (s *UploadService) Delete(uploadId int) error {
	return s.repo.Delete(uploadId)
}
