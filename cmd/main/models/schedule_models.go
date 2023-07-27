package models

type Physicians struct {
	FirstName string `json:"physician_first_name"`
	LastName  string `json:"physician_last_name"`
}

type PhysicianDataList struct {
	PhysicianList []PhysicanData
}

type PhysicanData struct {
	PhysicanID int    `json:"physican_id"`
	FirstName  string `json:"physician_first_name"`
	LastName   string `json:"physician_last_name"`
}
