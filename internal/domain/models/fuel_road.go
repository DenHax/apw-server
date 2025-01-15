package models

type FuelRoad struct {
	FuelRoadNumber int    `db:"fuel_road_number" json:"fuel_road_number"`
	TypeId         int    `db:"fuel_type_id" json:"type_id"`
	TypeName       string `db:"type_name" json:"type_name"`
	Mass           int    `db:"fuel_road_mass" json:"mass"`
	Condition      string `db:"fuel_road_condition" json:"condition"`
}
