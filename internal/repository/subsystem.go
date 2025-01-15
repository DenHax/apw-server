package repo

import (
	"apw/internal/domain/models"
	"apw/internal/storage"
	"fmt"
)

type SubsystemPsql struct {
	storage *storage.Storage
}

func NewSubsystemPostgres(s *storage.Storage) *SubsystemPsql {
	return &SubsystemPsql{storage: s}
}

func (r *SubsystemPsql) GetAll() ([]models.Subsystem, error) {
	var subsystems []models.Subsystem

	query := fmt.Sprintf(`SELECT ss.subsystem_id, ss.subsystem_name, ss.subsystem_status
  FROM %s AS ss WHERE ss.subsystem_name = 'реактор'`, subsystemTable)
	err := r.storage.DB.Select(&subsystems, query)

	return subsystems, err
}
