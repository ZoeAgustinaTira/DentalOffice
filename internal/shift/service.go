package shift

import (
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
)

type Service interface {
	Save(s domain.Shift) (domain.Shift, error)
	GetByID(id int) (domain.Shift, error)
	Update(s domain.Shift) (domain.Shift, error)
	UpdateAll(s domain.Shift) (domain.Shift, error)
	Delete(id int) error
	GetByDNI() (domain.Shift, error)
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

func (s *service) Save(sh domain.Shift) (domain.Shift, error) {
	//newDentist := domain.n(d.Name, d.Surname, d.Enrollment)

	id, err := s.repository.Save(sh)
	if err != nil {
		return domain.Shift{}, err
	}

	sh.ID = id

	return sh, nil
}

func (s *service) GetByID(id int) (domain.Shift, error) {
	shift, err := s.repository.GetByID(id)
	if err != nil {
		return domain.Shift{}, err
	}

	return shift, nil

}

func (s *service) GetByDNI() (domain.Shift, error) {
	shift, err := s.repository.GetByDNI()
	if err != nil {
		return domain.Shift{}, err
	}

	return shift, nil
}

func (s *service) Update(sh domain.Shift) (domain.Shift, error) {
	shift, err := s.GetByID(sh.ID)
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

	shUpdate, err := s.repository.Update(shift)
	if err != nil {
		return domain.Shift{}, err
	}
	return shUpdate, nil
}

func (s *service) UpdateAll(sh domain.Shift) (domain.Shift, error) {
	shUpdate, err := s.repository.Update(sh)
	if err != nil {
		return domain.Shift{}, err
	}
	return shUpdate, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
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
