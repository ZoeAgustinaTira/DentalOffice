package domain

type Shift struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	DentistID int    `json:"dentist_id"`
	PatientID int    `json:"patient_id"`
}
