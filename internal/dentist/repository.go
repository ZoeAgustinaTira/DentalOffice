package dentist

import (
	"database/sql"
	"fmt"
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
)

type Repository interface {
	Save(d domain.Dentist) (int, error)
	GetByID(id int) (domain.Dentist, error)
	Update(d domain.Dentist) (domain.Dentist, error)
	//UpdateAll(d domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
	Exists(enrollment string) bool
}

const (
	SAVE_DENTIST         = "INSERT INTO dentists(name,surname,enrollment) VALUES (?,?,?);"
	GET_DENTIST_BY_ID    = "SELECT * FROM dentists WHERE id = ?;"
	UPDATE_DENTIST       = "UPDATE dentists SET name = ?, surname = ?, enrollment = ? WHERE id = ?;"
	DELETE_DENTIST_BY_ID = "DELETE FROM dentists WHERE id = ?;"
	EXIST_DENTIST        = "SELECT enrollment FROM dentists WHERE enrollment = ?"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Save(d domain.Dentist) (int, error) {
	stmt, err := r.db.Prepare(SAVE_DENTIST)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&d.Name, &d.Surname, &d.Enrollment)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *repository) GetByID(id int) (domain.Dentist, error) {
	row := r.db.QueryRow(GET_DENTIST_BY_ID, id)
	d := domain.Dentist{}
	err := row.Scan(&d.ID, &d.Name, &d.Surname, &d.Enrollment)
	if err != nil {
		return d, err
	}

	return d, nil
}

func (r *repository) Update(d domain.Dentist) (domain.Dentist, error) {
	stmt, err := r.db.Prepare(UPDATE_DENTIST)
	if err != nil {
		return domain.Dentist{}, err
	}

	res, err := stmt.Exec(&d.Name, &d.Surname, &d.Enrollment, &d.ID)
	if err != nil {
		return domain.Dentist{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}

	return d, nil
}

func (r *repository) UpdateAll(d domain.Dentist) (domain.Dentist, error) {
	stmt, err := r.db.Prepare(UPDATE_DENTIST)
	if err != nil {
		return domain.Dentist{}, err
	}

	res, err := stmt.Exec(&d.Name, &d.Surname, &d.Enrollment, &d.ID)
	if err != nil {
		return domain.Dentist{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}

	return d, nil
}

func (r *repository) Delete(id int) error {
	stmt, err := r.db.Prepare(DELETE_DENTIST_BY_ID)
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

func (r *repository) Exists(enrollment string) bool {
	row := r.db.QueryRow(EXIST_DENTIST, enrollment)
	err := row.Scan(&enrollment)
	return err == nil
}
