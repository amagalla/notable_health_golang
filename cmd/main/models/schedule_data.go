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
