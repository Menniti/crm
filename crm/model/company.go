package model

//Company collects data from company
type Company struct {
	Name          string `json:"name"`
	Contacts      string `json:"contacts"`
	RelationOwner string `json:"relationOwner"`
	Location      string `json:"location"`
	Status        string `json:"status"`
}
