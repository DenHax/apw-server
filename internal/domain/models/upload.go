package models

import "time"

type Upload struct {
	LoadDate         time.Time `db:"load_date" json:"load_date"`
	EmployeeId       *int      `db:"employee_id" json:"employee_id"`
	EmployeeFullName string    `json:"employee_full_name" db:"employee_full_name"`
	SubsystemId      *int      `db:"subsystem_id" json:"subsystem_id"`
	FuelRoadNumber   int       `db:"fuel_road_number" json:"fuel_road_number"`
}

type Report struct {
	EmployeeFullName string    `json:"employee_full_name" db:"employee_full_name"`
	SubsystemId      int       `json:"subsystem_id" db:"subsystem_id"`
	LoadCount        int       `json:"load_count" db:"load_count"`
	FirstLoadDate    time.Time `json:"first_load_date" db:"first_load_date"`
	LastLoadDate     time.Time `json:"last_load_date" db:"last_load_date"`
}
