package domain

type Patient struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Address       string `json:"address"`
	DNI           string `json:"DNI"`
	DischargeDate string `json:"discharge_date"`
}

func NewPatient(name string, surname string, address string, dni string, dischargeDate string) *Patient {
	return &Patient{
		ID:            0,
		Name:          name,
		Surname:       surname,
		Address:       address,
		DNI:           dni,
		DischargeDate: dischargeDate,
	}
}
