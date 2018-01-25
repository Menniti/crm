package model

//Client collects data from Client
type Client struct {
	Name          string `json:"name"`
	ID            int    `json:"id"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	ClientOwnerID int    `json:"clientOwnerID"`
	CompanyID     int    `json:"companyID"`
}
