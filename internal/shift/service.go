package shift

import (
	"context"
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
)

type Service interface {
	Save(ctx context.Context, s domain.Shift) (domain.Shift, error)
	GetByID(ctx context.Context, id int) (domain.Shift, error)
	Update(ctx context.Context, s domain.Shift) (domain.Shift, error)
	Delete(ctx context.Context, id int) error
	GetByDNI(ctx context.Context) (domain.Shift, error)
	//Exists(ctx context.Context, enrollment string) bool
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Save(ctx context.Context, sh domain.Shift) (domain.Shift, error) {
	//newDentist := domain.n(d.Name, d.Surname, d.Enrollment)

	id, err := s.repository.Save(ctx, sh)
	if err != nil {
		return domain.Shift{}, err
	}

	sh.ID = id

	return sh, nil
}

func (s *service) GetByID(ctx context.Context, id int) (domain.Shift, error) {
	shift, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return domain.Shift{}, err
	}

	return shift, nil

}

func (s *service) GetByDNI(ctx context.Context) (domain.Shift, error) {
	shift, err := s.repository.GetByDNI(ctx)
	if err != nil {
		return domain.Shift{}, err
	}

	return shift, nil
}

func (s *service) Update(ctx context.Context, sh domain.Shift) (domain.Shift, error) {
	shift, err := s.GetByID(ctx, sh.ID)
	if err != nil {
		return domain.Shift{}, err
	}

	if sh.Date == "" {
		sh.Date = shift.Date
	}
	if sh.Time == "" {
		sh.Time = shift.Time
	}
	if sh.DentistID == 0 {
		sh.DentistID = shift.DentistID
	}
	if sh.PatientID == 0 {
		sh.PatientID = shift.PatientID
	}

	//	dentistToUpdate := domain.NewDentist(d.Name, d.Surname, d.Enrollment)
	//dentistToUpdate.ID = d.ID

	shUpdate, err := s.repository.Update(ctx, shift)
	if err != nil {
		return domain.Shift{}, err
	}
	return shUpdate, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

/*
func (s *service) Exists(ctx context.Context, enrollment string) bool {
	exist := s.repository.Exists(ctx, enrollment)
	return exist
}
*/
