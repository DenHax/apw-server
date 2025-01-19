package repo

import (
	"apw/internal/domain/models"
	"apw/internal/storage"
	"fmt"
	"time"

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

	query := fmt.Sprintf(`SELECT up.load_date, up.subsystem_id, up.employee_id, up.fuel_road_number, em.employee_surname || ' ' || em.employee_firstname || 
    COALESCE(' ' || employee_lastname, '') AS employee_full_name FROM %s AS up JOIN %s AS em ON up.employee_id = em.employee_id
    JOIN subsystem AS ss ON ss.subsystem_id = up.subsystem_id`,
		uploadTable, employeeTable)
	err := r.storage.DB.Select(&uploads, query)
	return uploads, err
}

func (r *UploadPsql) GetReport(fdate time.Time, sdate time.Time) ([]models.Report, error) {
	var reports []models.Report
	fmt.Println(fdate, sdate)

	query := fmt.Sprintf(` SELECT 
        e.employee_surname || ' ' || e.employee_firstname as employee_full_name,
        e.subsystem_id,
        COUNT(u.load_date) AS load_count,
        MIN(u.load_date) AS first_load_date,
        MAX(u.load_date) AS last_load_date
      FROM %s u
      JOIN %s e ON u.employee_id = e.employee_id
	  JOIn %s ss on u.subsystem_id = ss.subsystem_id
      WHERE u.load_date BETWEEN $1 AND $2
      GROUP BY e.employee_id, employee_full_name, e.subsystem_id`,
		uploadTable, employeeTable, subsystemTable)
	err := r.storage.DB.Select(&reports, query, fdate, sdate)
	fmt.Println(err)
	return reports, err
}

func (r *UploadPsql) Create(upload models.Upload) (uploadId time.Time, err error) {
	fmt.Println("create upload")
	var id time.Time

	var subsystemId interface{} = nil
	if upload.SubsystemId != nil {
		subsystemId = *upload.SubsystemId
	}

	var employeeId interface{} = nil
	if upload.EmployeeId != nil {
		employeeId = *upload.EmployeeId
	}

	fmt.Println(subsystemId, employeeId, upload.LoadDate)
	query := fmt.Sprintf(`
        INSERT INTO %s (load_date, subsystem_id, employee_id, fuel_road_number)
        VALUES ($1, $2, $3, $4)
        RETURNING load_date
    `, uploadTable)

	row := r.storage.DB.QueryRow(query, upload.LoadDate, subsystemId, employeeId, upload.FuelRoadNumber)
	if err := row.Scan(&id); err != nil {
		logrus.Errorf("Failed to create upload: %v", err)
		return id, err
	}

	return id, nil
}

func (r *UploadPsql) Delete(uploadId time.Time) (err error) {
	query := fmt.Sprintf(`DELETE FROM %s upload 
									WHERE load_date = $1`, uploadTable)
	status, err := r.storage.DB.Exec(query, uploadId)
	logrus.Debugf("/api/upload/:id delete: %s", status)
	return err
}
