package model

import "gopkg.in/mgo.v2/bson"

//Company collects data from company
type Company struct {
	Name     string        `json:"name" bson:"name"`
	Contacts bson.ObjectId `json:"contacts" bson:"contacts,omitempty"`
	OwnerID  bson.ObjectId `json:"ownerID" bson:"ownerID,omitempty"`
	Location string        `json:"location" bson:"location"`
	Status   string        `json:"status" bson:"status"`
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
}
