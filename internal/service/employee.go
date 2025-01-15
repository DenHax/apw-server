package service

import (
	"apw/internal/domain/models"
	repo "apw/internal/repository"
)

type EmployeeService struct {
	repo repo.Employee
}

func NewEmployeeService(repo repo.Employee) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) GetAll() ([]models.Employee, error) {
	return s.repo.GetAll()
}
