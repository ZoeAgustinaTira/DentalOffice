package patient

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, p domain.Patient) (int, error)
	GetByID(ctx context.Context, id int) (domain.Patient, error)
	Update(ctx context.Context, p domain.Patient) (domain.Patient, error)
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, patientID string) bool
}

const (
	SAVE_PATIENT         = "INSERT INTO dentaloffice.patients(name, surname, address,DNI,dischargeDate) VALUES (?,?,?,?,?);"
	GET_PATIENT_BY_ID    = "SELECT * FROM patients WHERE id = ?;"
	UPDATE_PATIENT       = "UPDATE dentaloffice.patients SET name = ?, surname = ?, address = ?, DNI = ?, dischargeDate = ? WHERE id = ?;"
	DELETE_PATIENT_BY_ID = "DELETE FROM dentaloffice.patients WHERE id = ?;"
	EXIST_PATIENT        = "SELECT id FROM dentaloffice.patients WHERE id = ?"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, p domain.Patient) (int, error) {
	stmt, err := r.db.Prepare(SAVE_PATIENT)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&p.Name, &p.Surname, &p.ID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *repository) GetByID(ctx context.Context, id int) (domain.Patient, error) {
	row := r.db.QueryRow(GET_PATIENT_BY_ID)
	p := domain.Patient{}
	err := row.Scan(&p.ID, &p.Name, &p.Surname, &p.ID)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (r *repository) Update(ctx context.Context, p domain.Patient) (domain.Patient, error) {
	stmt, err := r.db.Prepare(UPDATE_PATIENT)
	if err != nil {
		return domain.Patient{}, err
	}

	res, err := stmt.Exec(&p.Name, &p.Surname, &p.ID)
	if err != nil {
		return domain.Patient{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}

	return p, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DELETE_PATIENT_BY_ID)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowAffect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowAffect < 1 {
		return fmt.Errorf("NotFound")
	}

	return nil
}

func (r *repository) Exists(ctx context.Context, patientID string) bool {
	row := r.db.QueryRow(EXIST_PATIENT, patientID)
	err := row.Scan(&patientID)
	return err == nil
}
