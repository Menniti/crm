package model

import "gopkg.in/mgo.v2/bson"

//Client collects data from Client
type Client struct {
	Name      string        `json:"name" bson:"name"`
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email     string        `json:"email" bson:"email"`
	Phone     string        `json:"phone" bson:"phone"`
	OwnerID   bson.ObjectId `json:"ownerID" bson:"_id,omitempty"`
	CompanyID bson.ObjectId `json:"companyID" bson:"_id,omitempty"`
}
