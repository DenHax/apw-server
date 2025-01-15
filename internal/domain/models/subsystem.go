package models

type Subsystem struct {
	SubsystemNumber int    `db:"subsystem_id" json:"subsystem_id"`
	Name            string `db:"subsystem_name" json:"name"`
	Status          string `db:"subsystem_status" json:"status"`
}
