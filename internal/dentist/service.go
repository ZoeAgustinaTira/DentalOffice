package dentist

import (
	"context"
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
)

type Service interface {
	Save(ctx context.Context, d domain.Dentist) (domain.Dentist, error)
	GetByID(ctx context.Context, id int) (domain.Dentist, error)
	Update(ctx context.Context, d domain.Dentist) (domain.Dentist, error)
	UpdateAll(ctx context.Context, d domain.Dentist) (domain.Dentist, error)
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, enrollment string) bool
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Save(ctx context.Context, d domain.Dentist) (domain.Dentist, error) {
	newDentist := domain.NewDentist(d.Name, d.Surname, d.Enrollment)

	id, err := s.repository.Save(ctx, *newDentist)
	if err != nil {
		return domain.Dentist{}, err
	}

	newDentist.ID = id

	return *newDentist, nil
}

func (s *service) GetByID(ctx context.Context, id int) (domain.Dentist, error) {
	dentist, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil

}

func (s *service) Update(ctx context.Context, d domain.Dentist) (domain.Dentist, error) {
	dentist, err := s.GetByID(ctx, d.ID)
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

	dUpdate, err := s.repository.Update(ctx, *dentistToUpdate)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dUpdate, nil
}

func (s *service) UpdateAll(ctx context.Context, d domain.Dentist) (domain.Dentist, error) {
	dUpdate, err := s.repository.Update(ctx, d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dUpdate, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Exists(ctx context.Context, enrollment string) bool {
	exist := s.repository.Exists(ctx, enrollment)
	return exist
}
