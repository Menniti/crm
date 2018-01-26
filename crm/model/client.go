package model

import "gopkg.in/mgo.v2/bson"

//Client collects data from Client
type Client struct {
	Name      string        `json:"name"`
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email     string        `json:"email" bson:"_id,omitempty"`
	Phone     string        `json:"phone"`
	OwnerID   bson.ObjectId `json:"ownerID" bson:"_id,omitempty"`
	CompanyID bson.ObjectId `json:"companyID" bson:"_id,omitempty"`
}
