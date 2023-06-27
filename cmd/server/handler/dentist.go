package handler

import "github.com/ZoeAgustinatira/DentalOffice/internal/dentist"

type Dentist struct {
	dentistService dentist.Service
}

func NewDentist(d dentist.Service) *Dentist {
	return &Dentist{
		dentistService: d,
	}
}
