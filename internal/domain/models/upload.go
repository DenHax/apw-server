package models

import "time"

type Upload struct {
	LoadDate         time.Time `db:"load_date" json:"load_date"`
	EmployeeId       int       `db:"employee_id" json:"employee_id"`
	EmployeeFullName string    `json:"employee_full_name" db:"employee_full_name"`
	SubsystemId      *int      `db:"subsystem_id" json:"subsystem_id"`
	SubsystemName    *string   `json:"subsystem_name" db:"subsystem_name"`
	FuelRoadNumber   int       `db:"fuel_road_number" json:"fuel_road_number"`
}
