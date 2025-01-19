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
func (s *EmployeeService) Update(employeeId int, input models.UpdateEmployee) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(employeeId, input)
}
