package models

type PhysicianListData struct {
	IdPhysician int    `json:"IdPhysician"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
}

type InsertPhysicianData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type AddAppointmentData struct {
	PatientFirstName string `json:"patientFirstName"`
	PatientLastName  string `json:"patientLastName"`
	Date_Column      string `json:"date_column"`
	Time_Column      string `json:"time_column"`
	Kind             string `json:"kind"`
}

type CheckAppointmentData struct {
	PatientLastName string `json:"patientLastName"`
}
