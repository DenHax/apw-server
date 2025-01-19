package repo

import (
	"apw/internal/domain/models"
	"apw/internal/storage"
	"fmt"
)

type FuelRoadPsql struct {
	storage *storage.Storage
}

func NewFuelRoadPostgres(s *storage.Storage) *FuelRoadPsql {
	return &FuelRoadPsql{storage: s}
}

func (r *FuelRoadPsql) GetAll() ([]models.FuelRoad, error) {
	var fuelRoads []models.FuelRoad

	query := fmt.Sprintf(`SELECT fr.fuel_road_number, fr.fuel_type_id, fr.fuel_road_mass, fr.fuel_road_condition,
ft.FUEL_TYPE_SHELL || '+' || ft.FUEL_TYPE_CONTACT || '+' || ft.FUEL_TYPE_FORM AS type_name
	   FROM %s AS fr JOIN %s AS ft ON fr.fuel_type_id = ft.fuel_type_id `, fuelRoadTable, fuelTypeTable)
	err := r.storage.DB.Select(&fuelRoads, query)
	return fuelRoads, err
}
