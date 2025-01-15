package repo

import (
	"apw/internal/domain/models"
	"apw/internal/storage"
	"fmt"
)

type EmployeePsql struct {
	storage *storage.Storage
}

func NewEmployeePostgres(s *storage.Storage) *EmployeePsql {
	return &EmployeePsql{storage: s}
}

func (r *EmployeePsql) GetAll() ([]models.Employee, error) {
	var employees []models.Employee

	query := fmt.Sprintf(`SELECT em.employee_firstname, em.employee_surname, em.employee_lastname, em.employee_title, em.employee_phone, em.subsystem_id 
	   FROM %s AS em`, employeeTable)
	err := r.storage.DB.Select(&employees, query)
	return employees, err
}
