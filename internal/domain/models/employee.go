package models

type Employee struct {
	EmpoyeeId   *int    `json:"employee_id" db:"employee_id"`
	Firstname   string  `json:"firstname" db:"employee_firstname"`
	Surname     string  `json:"surname" db:"employee_surname"`
	Lastname    *string `json:"lastname" db:"employee_lastname"`
	SubsystemId int     `json:"subsystem_id" db:"subsystem_id"`
	Title       string  `db:"employee_title" json:"title"`
	Phone       string  `db:"employee_phone" json:"phone"`
}
