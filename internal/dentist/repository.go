package dentist

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
)

type Repository interface {
	Save(ctx context.Context, d domain.Dentist) (int, error)
	GetByID(ctx context.Context, id int) (domain.Dentist, error)
	Update(ctx context.Context, d domain.Dentist) (domain.Dentist, error)
	//UpdateAll(ctx context.Context, d domain.Dentist) (domain.Dentist, error)
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, enrollment string) bool
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

func (r *repository) Save(ctx context.Context, d domain.Dentist) (int, error) {
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

func (r *repository) GetByID(ctx context.Context, id int) (domain.Dentist, error) {
	row := r.db.QueryRow(GET_DENTIST_BY_ID, id)
	d := domain.Dentist{}
	err := row.Scan(&d.ID, &d.Name, &d.Surname, &d.Enrollment)
	if err != nil {
		return d, err
	}

	return d, nil
}

func (r *repository) Update(ctx context.Context, d domain.Dentist) (domain.Dentist, error) {
	stmt, err := r.db.Prepare(UPDATE_DENTIST)
	if err != nil {
		return domain.Dentist{}, err
	}

	res, err := stmt.Exec(&d.ID, &d.Name, &d.Surname, &d.Enrollment)
	if err != nil {
		return domain.Dentist{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}

	return d, nil
}

func (r *repository) UpdateAll(ctx context.Context, d domain.Dentist) (domain.Dentist, error) {
	stmt, err := r.db.Prepare(UPDATE_DENTIST)
	if err != nil {
		return domain.Dentist{}, err
	}

	res, err := stmt.Exec(&d.Name, &d.Surname, &d.Enrollment)
	if err != nil {
		return domain.Dentist{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}

	return d, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
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

func (r *repository) Exists(ctx context.Context, enrollment string) bool {
	row := r.db.QueryRow(EXIST_DENTIST, enrollment)
	err := row.Scan(&enrollment)
	return err == nil
}
