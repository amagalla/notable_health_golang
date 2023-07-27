package models

type Physicians struct {
	FirstName string `json:"physician_first_name"`
	LastName  string `json:"physician_last_name"`
}

type PhysicianResponse struct {
	PhysicianList []PhysicianData
}

type PhysicianData struct {
	PhysicianID int    `json:"physician_id"`
	FirstName   string `json:"physician_first_name"`
	LastName    string `json:"physician_last_name"`
}
