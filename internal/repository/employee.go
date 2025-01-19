package repo

import (
	"apw/internal/domain/models"
	"apw/internal/storage"
	"fmt"
	"strings"
)

type EmployeePsql struct {
	storage *storage.Storage
}

func NewEmployeePostgres(s *storage.Storage) *EmployeePsql {
	return &EmployeePsql{storage: s}
}

func (r *EmployeePsql) GetAll() ([]models.Employee, error) {
	var employees []models.Employee

	query := fmt.Sprintf(`SELECT em.employee_id, em.employee_firstname, em.employee_surname, em.employee_lastname,
     em.employee_title, em.employee_phone, em.subsystem_id 
	   FROM %s AS em WHERE em.employee_title LIKE '%%реактор%%'`, employeeTable)
	err := r.storage.DB.Select(&employees, query)
	return employees, err
}

func (r *EmployeePsql) Update(employeeId int, input models.UpdateEmployee) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Firstname != nil {
		setValues = append(setValues, fmt.Sprintf("employee_firstname=$%d", argId))
		args = append(args, *input.Firstname)
		argId++
	}

	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("employee_surname=$%d", argId))
		args = append(args, *input.Surname)
		argId++
	}

	if input.Lastname != nil {
		setValues = append(setValues, fmt.Sprintf("employee_lastname=$%d", argId))
		args = append(args, *input.Lastname)
		argId++
	}

	if input.SubsystemId != nil {
		setValues = append(setValues, fmt.Sprintf("subsystem_id=$%d", argId))
		args = append(args, *input.SubsystemId)
		argId++
	}
	if input.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("employee_phone=$%d", argId))
		args = append(args, *input.Phone)
		argId++
	}
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s em SET %s WHERE em.employee_id = $%d",
		employeeTable, setQuery, argId)
	args = append(args, employeeId)

	fmt.Println("updateQuery: %s", query)
	fmt.Println("args: %s", args)

	_, err := r.storage.DB.Exec(query, args...)
	return err
}
