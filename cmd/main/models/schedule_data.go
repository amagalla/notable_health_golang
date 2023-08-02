package models

type PostPhysiciansData struct {
	FirstName string `json:"physician_first_name"`
	LastName  string `json:"physician_last_name"`
}

type GetPhysicanData struct {
	PhysicanID int    `json:"physican_id"`
	FirstName  string `json:"physician_first_name"`
	LastName   string `json:"physician_last_name"`
}

type PostAppointmentData struct {
	PatientFirstName string `json:"patient_first_name"`
	PatientLastName  string `json:"patient_last_name"`
	ScheduledDate    string `json:"scheduled_date"`
	ScheduledTime    string `json:"scheduled_time"`
	Kind             string `json:"kind"`
	PhysicianID      uint64 `json:"physician_id"`
}
