package repo

import (
	"apw/internal/domain/models"
	"apw/internal/storage"
	"fmt"
)

type FuelTypePsql struct {
	storage *storage.Storage
}

func NewFuelTypePostgres(s *storage.Storage) *FuelTypePsql {
	return &FuelTypePsql{storage: s}
}

func (r *FuelTypePsql) GetAll() ([]models.FuelType, error) {
	var fuelTypes []models.FuelType

	query := fmt.Sprintf(`SELECT ft.fuel_type_id, ft.fuel_type_shell, ft.fuel_type_form, ft.fuel_type_contact
    FROM %s AS ft`, fuelTypeTable)
	err := r.storage.DB.Select(&fuelTypes, query)

	return fuelTypes, err
}
