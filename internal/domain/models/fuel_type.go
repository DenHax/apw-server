package models

type FuelType struct {
	FuelTypeId int    `json:"fuel_type_id" db:"fuel_type_id"`
	Shell      string `json:"shell" db:"fuel_type_shell"`
	Contact    string `json:"contact" db:"fuel_type_contact"`
	Form       string `json:"form" db:"fuel_type_form"`
}
