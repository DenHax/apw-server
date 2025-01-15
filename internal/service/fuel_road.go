package service

import (
	"apw/internal/domain/models"
	repo "apw/internal/repository"
)

type FuelRoadService struct {
	repo repo.FuelRoad
}

func NewFuelRoadService(repo repo.FuelRoad) *FuelRoadService {
	return &FuelRoadService{repo: repo}
}

func (s *FuelRoadService) GetAll() ([]models.FuelRoad, error) {
	return s.repo.GetAll()
}
