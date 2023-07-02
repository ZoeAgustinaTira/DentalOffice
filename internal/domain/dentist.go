package domain

type Dentist struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Enrollment string `json:"enrollment"`
}

func NewDentist(n string, sn string, e string) *Dentist {
	return &Dentist{
		ID:         0,
		Name:       n,
		Surname:    sn,
		Enrollment: e,
	}
}
