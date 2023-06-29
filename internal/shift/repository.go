package shift

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, s domain.Shift) (int, error)
	GetByID(ctx context.Context, id int) (domain.Shift, error)
	Update(ctx context.Context, s domain.Shift) (domain.Shift, error)
	Delete(ctx context.Context, id int) error
	GetByDNI(ctx context.Context) (domain.Shift, error)
	//Exists(ctx context.Context, enrollment string) bool
}

const (
	SAVE_SHIFT         = "INSERT INTO shifts(date, time, dentist_id, patient_id) VALUES (?, ?, ?, ?);"
	GET_SHIFT_BY_ID    = "SELECT * FROM shifts WHERE id = ?;"
	GET_SHIFT_BY_DNI   = "SELECT * FROM shifts s JOIN patients p ON p.id = s.patient_id where p.dni = ? GROUP BY p.dni;"
	UPDATE_SHIFT       = "UPDATE shifts SET date = ?, time = ?, dentist_id = ?, patient_id = ? WHERE id = ?;"
	DELETE_SHIFT_BY_ID = "DELETE FROM shifts WHERE id = ?;"
	EXIST_SHIFT        = "SELECT enrollment FROM shifts WHERE enrollment = ?"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(ctx context.Context, s domain.Shift) (int, error) {
	stmt, err := r.db.Prepare(SAVE_SHIFT)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&s.Date, &s.Time, &s.DentistID, &s.PatientID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *repository) GetByID(ctx context.Context, id int) (domain.Shift, error) {
	row := r.db.QueryRow(GET_SHIFT_BY_ID, id)
	s := domain.Shift{}
	err := row.Scan(&s.ID, &s.Date, &s.Time, &s.DentistID, s.PatientID)
	if err != nil {
		return s, err
	}

	return s, nil
}

func (r *repository) GetByDNI(ctx context.Context) (domain.Shift, error) {
	row := r.db.QueryRow(GET_SHIFT_BY_DNI)
	s := domain.Shift{}
	err := row.Scan(&s.ID, &s.Date, &s.Time, &s.DentistID, s.PatientID)
	if err != nil {
		return s, err
	}

	return s, nil
}

func (r *repository) Update(ctx context.Context, s domain.Shift) (domain.Shift, error) {
	stmt, err := r.db.Prepare(UPDATE_SHIFT)
	if err != nil {
		return domain.Shift{}, err
	}

	res, err := stmt.Exec(&s.Date, &s.Time, &s.DentistID, &s.PatientID)
	if err != nil {
		return domain.Shift{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Shift{}, err
	}

	return s, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(DELETE_SHIFT_BY_ID)
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

/*
func (r *repository) Exists(ctx context.Context, enrollment string) bool {
	row := r.db.QueryRow(EXIST_DENTIST, enrollment)
	err := row.Scan(&enrollment)
	return err == nil
}
*/
