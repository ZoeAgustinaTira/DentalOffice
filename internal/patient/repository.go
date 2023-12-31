package patient

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, p domain.Patient) (int, error)
	GetByID(id int) (domain.Patient, error)
	Update(p domain.Patient) (domain.Patient, error)
	Delete(ctx context.Context, id int) error
	Exists(DNI string) bool
	HasShifts(id int) bool
}

const (
	SAVE_PATIENT         = "INSERT INTO patients(name, surname, address,DNI,dischargeDate) VALUES (?,?,?,?,?);"
	GET_PATIENT_BY_ID    = "SELECT * FROM patients WHERE id = ?;"
	UPDATE_PATIENT       = "UPDATE patients SET name = ?, surname = ?, address = ?, DNI = ?, dischargeDate = ? WHERE id = ?;"
	DELETE_PATIENT_BY_ID = "DELETE FROM patients WHERE id = ?;"
	EXIST_PATIENT        = "SELECT DNI FROM patients WHERE DNI = ?"
	HAS_SHIFT            = "SELECT p.* FROM patients p INNER JOIN shifts s ON p.id = s.patient_id WHERE patient_id = ?;"
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

	res, err := stmt.Exec(&p.Name, &p.Surname, &p.Address, &p.DNI, &p.DischargeDate)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *repository) GetByID(id int) (domain.Patient, error) {
	row := r.db.QueryRow(GET_PATIENT_BY_ID, id)
	p := domain.Patient{}
	err := row.Scan(&p.ID, &p.Name, &p.Surname, &p.Address, &p.DNI, &p.DischargeDate)
	if err != nil {
		return p, err
	}

	return p, nil
}

func (r *repository) Update(p domain.Patient) (domain.Patient, error) {
	stmt, err := r.db.Prepare(UPDATE_PATIENT) //Asegúrate de tener una constante SQL apropiada para esto
	if err != nil {
		return domain.Patient{}, err
	}

	res, err := stmt.Exec(&p.Name, &p.Surname, &p.Address, &p.DNI, &p.DischargeDate, &p.ID)
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

func (r *repository) Exists(DNI string) bool {
	row := r.db.QueryRow(EXIST_PATIENT, DNI)
	err := row.Scan(&DNI)
	return err == nil
}

func (r *repository) HasShifts(id int) bool {
	rows, _ := r.db.Query(HAS_SHIFT, id)

	if rows != nil {
		return true
	}
	return false
}
