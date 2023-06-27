package domain

type Patient struct {
	Name          string `json:"name"`
	Surname       string `json:"surname"`
	Address       string `json:"address"`
	DNI           string `json:"DNI"`
	DischargeDate string `json:"discharge_date"`
}
