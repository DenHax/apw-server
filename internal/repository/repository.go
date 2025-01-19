package repo

import (
	"apw/internal/domain/models"
	"apw/internal/storage"
	"time"
)

const (
	uploadTable    = "upload"
	subsystemTable = "subsystem"
	employeeTable  = "employee"
	fuelRoadTable  = "fuel_road"
	fuelTypeTable  = "fuel_type"
)

type Employee interface {
	GetAll() ([]models.Employee, error)
	Update(int, models.UpdateEmployee) error
}

type Upload interface {
	GetAll() ([]models.Upload, error)
	GetReport(time.Time, time.Time) ([]models.Report, error)
	Create(upload models.Upload) (time.Time, error)
	Delete(uploadId time.Time) error
}
type FuelRoad interface {
	GetAll() ([]models.FuelRoad, error)
}
type FuelType interface {
	GetAll() ([]models.FuelType, error)
}
type Subsystem interface {
	GetAll() ([]models.Subsystem, error)
}

type Repository struct {
	Employee
	Upload
	FuelRoad
	FuelType
	Subsystem
}

func NewRepository(s *storage.Storage) *Repository {
	return &Repository{
		Employee:  NewEmployeePostgres(s),
		Subsystem: NewSubsystemPostgres(s),
		FuelRoad:  NewFuelRoadPostgres(s),
		FuelType:  NewFuelTypePostgres(s),
		Upload:    NewUploadPostgres(s),
	}
}
