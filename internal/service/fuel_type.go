package service

import (
	"apw/internal/domain/models"
	repo "apw/internal/repository"
)

type FuelTypeService struct {
	repo repo.FuelType
}

func NewFuelTypeService(repo repo.FuelType) *FuelTypeService {
	return &FuelTypeService{repo: repo}
}

func (s *FuelTypeService) GetAll() ([]models.FuelType, error) {
	return s.repo.GetAll()
}
