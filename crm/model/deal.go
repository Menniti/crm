package model

//Deal collects data from Deal
type Deal struct {
	Name      string `json:"name"`
	DealOwner string `json:"dealOwner"`
	Status    string `json:"status"`
	Value     int    `json:"value"`
}
