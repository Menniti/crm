package model

import "gopkg.in/mgo.v2/bson"

//Client collects data from Client
type Client struct {
	Name          string        `json:"name"`
	ID            bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email         string        `json:"email"`
	Phone         string        `json:"phone"`
	ClientOwnerID int           `json:"clientOwnerID"`
	CompanyID     int           `json:"companyID"`
}
