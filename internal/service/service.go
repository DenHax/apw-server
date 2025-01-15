package service

import (
	"apw/internal/domain/models"
	repo "apw/internal/repository"
)

type Employee interface {
	GetAll() ([]models.Employee, error)
}

type Upload interface {
	GetAll() ([]models.Upload, error)
	Create(upload models.Upload) (int, error)
	Delete(uploadId int) error
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

type Service struct {
	Employee
	Upload
	FuelType
	FuelRoad
	Subsystem
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		Employee:  NewEmployeeService(repos.Employee),
		FuelRoad:  NewFuelRoadService(repos.FuelRoad),
		FuelType:  NewFuelTypeService(repos.FuelType),
		Subsystem: NewSubsystemService(repos.Subsystem),
		Upload:    NewUploadService(repos.Upload),
	}
}
