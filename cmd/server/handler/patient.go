package handler

import (
	"github.com/ZoeAgustinatira/DentalOffice/internal/patient"
)

type Patient struct {
	patientService patient.Service
}

func NewPatient(p patient.Service) *Patient {
	return &Patient{
		patientService: p,
	}
}
