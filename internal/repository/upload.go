package repo

import (
	"apw/internal/domain/models"
	"apw/internal/storage"
	"fmt"

	"github.com/sirupsen/logrus"
)

type UploadPsql struct {
	storage *storage.Storage
}

func NewUploadPostgres(s *storage.Storage) *UploadPsql {
	return &UploadPsql{storage: s}
}

func (r *UploadPsql) GetAll() ([]models.Upload, error) {
	var uploads []models.Upload

	query := fmt.Sprintf(`SELECT up.load_date, up.subsystem_id, up.fuel_road_number, em.employee_surname || ' ' || em.employee_firstname || 
    COALESCE(' ' || employee_lastname, '') employee_full_name FROM %s AS up JOIN %s AS em ON up.employee_id = em.employee_id`,
		uploadTable, employeeTable)
	err := r.storage.DB.Select(&uploads, query)

	return uploads, err
}

func (r *UploadPsql) Create(upload models.Upload) (uploadId int, err error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (load_date, subsystem_id, employee_id, fuel_road_number) values ($1, $2, $3, $4) RETURNING id`, uploadTable)

	row := r.storage.DB.QueryRow(query, upload.LoadDate, upload.SubsystemId, upload.EmployeeId, upload.FuelRoadNumber)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UploadPsql) Delete(uploadId int) (err error) {
	query := fmt.Sprintf(`DELETE FROM %s upload AS up
									WHERE up.load_date = $1`, uploadTable)
	status, err := r.storage.DB.Exec(query, uploadId)
	logrus.Debugf("/api/upload/:id delete: %s", status)
	return err
}
