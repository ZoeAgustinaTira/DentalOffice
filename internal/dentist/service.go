package dentist

import (
	"errors"
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
)

type Service interface {
	Save(d domain.Dentist) (domain.Dentist, error)
	GetByID(id int) (domain.Dentist, error)
	Update(d domain.Dentist) (domain.Dentist, error)
	UpdateAll(d domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
	Exists(enrollment string) bool
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Save(d domain.Dentist) (domain.Dentist, error) {
	newDentist := domain.NewDentist(d.Name, d.Surname, d.Enrollment)

	id, err := s.repository.Save(*newDentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	newDentist.ID = id

	return *newDentist, nil
}

func (s *service) GetByID(id int) (domain.Dentist, error) {
	dentist, err := s.repository.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil

}

func (s *service) Update(d domain.Dentist) (domain.Dentist, error) {
	dentist, err := s.GetByID(d.ID)
	if err != nil {
		return domain.Dentist{}, err
	}

	if d.Name == "" {
		d.Name = dentist.Name
	}
	if d.Surname == "" {
		d.Surname = dentist.Surname
	}
	if d.Enrollment == "" {
		d.Enrollment = dentist.Enrollment
	}

	dentistToUpdate := domain.NewDentist(d.Name, d.Surname, d.Enrollment)
	dentistToUpdate.ID = d.ID

	dUpdate, err := s.repository.Update(*dentistToUpdate)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dUpdate, nil
}

func (s *service) UpdateAll(d domain.Dentist) (domain.Dentist, error) {
	dUpdate, err := s.repository.Update(d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dUpdate, nil
}

func (s *service) Delete(id int) error {
	hasShift := s.repository.HasShifts(id)
	if hasShift {
		return errors.New("the dentist has an assigned shift")
	}

	err := s.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Exists(enrollment string) bool {
	exist := s.repository.Exists(enrollment)
	return exist
}
