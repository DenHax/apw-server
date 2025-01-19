package models

import "errors"

type Employee struct {
	EmpoyeeId   *int    `json:"employee_id" db:"employee_id"`
	Firstname   string  `json:"firstname" db:"employee_firstname"`
	Surname     string  `json:"surname" db:"employee_surname"`
	Lastname    *string `json:"lastname" db:"employee_lastname"`
	SubsystemId int     `json:"subsystem_id" db:"subsystem_id"`
	Title       string  `db:"employee_title" json:"title"`
	Phone       string  `db:"employee_phone" json:"phone"`
}

type UpdateEmployee struct {
	Firstname   *string `json:"firstname" db:"employee_firstname"`
	Surname     *string `json:"surname" db:"employee_surname"`
	Lastname    *string `json:"lastname" db:"employee_lastname"`
	SubsystemId *int    `json:"subsystem_id" db:"subsystem_id"`
	Phone       *string `json:"phone" db:"employee_phone"`
}

func (i UpdateEmployee) Validate() error {
	if i.Firstname == nil && i.Surname == nil && i.Lastname == nil && i.SubsystemId == nil && i.Phone == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
