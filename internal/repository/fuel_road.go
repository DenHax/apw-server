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

	query := fmt.Sprintf(`SELECT fr.fuel_road_number, fr.fuel_type_id, fr.fuel_road_mass, fr.fuel_road_condition
	   FROM %s AS fr`, fuelRoadTable)
	err := r.storage.DB.Select(&fuelRoads, query)
	return fuelRoads, err
}
