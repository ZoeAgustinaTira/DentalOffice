package patient

import (
	"context"
	"errors"
	"github.com/ZoeAgustinatira/DentalOffice/internal/domain"
)

type Service interface {
	Save(ctx context.Context, p domain.Patient) (domain.Patient, error)
	GetByID(id int) (domain.Patient, error)
	Update(p domain.Patient) (domain.Patient, error)
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

func (s *service) GetByID(id int) (domain.Patient, error) {
	patient, err := s.repository.GetByID(id)
	if err != nil {
		return domain.Patient{}, err
	}

	return patient, nil
}

func (s *service) Update(p domain.Patient) (domain.Patient, error) {
	patient, err := s.GetByID(p.ID)
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

	pUpdate, err := s.repository.Update(*patientToUpdate)
	if err != nil {
		return domain.Patient{}, err
	}
	return pUpdate, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	hasShift := s.repository.HasShifts()
	if hasShift {
		return errors.New("the patient has an assigned shift")
	}

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
