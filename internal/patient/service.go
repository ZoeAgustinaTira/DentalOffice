package patient

import (
	"context"
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
)

type Service interface {
	Save(ctx context.Context, p domain.Patient) (domain.Patient, error)
	GetByID(ctx context.Context, id int) (domain.Patient, error)
	Update(ctx context.Context, p domain.Patient) (domain.Patient, error)
	Delete(ctx context.Context, id int) error
	Exists(dni string) bool
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Save(ctx context.Context, p domain.Patient) (domain.Patient, error) {
	newPatient := domain.NewPatient(p.Name, p.Surname, p.Address, p.DNI, p.DischargeDate)

	id, err := s.repository.Save(ctx, *newPatient)
	if err != nil {
		return domain.Patient{}, err
	}

	newPatient.ID = id

	return *newPatient, nil
}

func (s *service) GetByID(ctx context.Context, id int) (domain.Patient, error) {
	patient, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return domain.Patient{}, err
	}

	return patient, nil
}

func (s *service) Update(ctx context.Context, p domain.Patient) (domain.Patient, error) {
	patient, err := s.GetByID(ctx, p.ID)
	if err != nil {
		return domain.Patient{}, err
	}

	if p.Name == "" {
		p.Name = patient.Name
	}
	if p.Surname == "" {
		p.Surname = patient.Surname
	}
	if p.Address == "" {
		p.Address = patient.Address
	}
	if p.DNI == "" {
		p.DNI = patient.DNI
	}
	if p.DischargeDate == "" {
		p.DischargeDate = patient.DischargeDate
	}

	patientToUpdate := domain.NewPatient(p.Name, p.Surname, p.Address, p.DNI, p.DischargeDate)
	patientToUpdate.ID = p.ID

	pUpdate, err := s.repository.Update(ctx, patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return pUpdate, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Exists(DNI string) bool {
	exist := s.repository.Exists(DNI)
	return exist
}
