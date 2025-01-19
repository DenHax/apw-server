package service

import (
	"apw/internal/domain/models"
	repo "apw/internal/repository"
	"time"
)

type UploadService struct {
	repo repo.Upload
}

func NewUploadService(repo repo.Upload) *UploadService {
	return &UploadService{repo: repo}
}

func (s *UploadService) Create(upload models.Upload) (time.Time, error) {
	return s.repo.Create(upload)
}

func (s *UploadService) GetAll() ([]models.Upload, error) {
	return s.repo.GetAll()
}
func (s *UploadService) GetReport(fdate, sdate time.Time) ([]models.Report, error) {
	return s.repo.GetReport(fdate, sdate)
}

func (s *UploadService) Delete(uploadId time.Time) error {
	return s.repo.Delete(uploadId)
}
