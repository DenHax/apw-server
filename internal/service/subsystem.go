package service

import (
	"apw/internal/domain/models"
	repo "apw/internal/repository"
)

type SubsystemService struct {
	repo repo.Subsystem
}

func NewSubsystemService(repo repo.Subsystem) *SubsystemService {
	return &SubsystemService{repo: repo}
}

func (s *SubsystemService) GetAll() ([]models.Subsystem, error) {
	return s.repo.GetAll()
}
