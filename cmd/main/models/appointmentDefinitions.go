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
