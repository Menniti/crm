package model

import "gopkg.in/mgo.v2/bson"

//Company collects data from company
type Company struct {
	Name     string        `json:"name"`
	Contacts bson.ObjectId `json:"contacts" bson:"_id,omitempty"`
	OwnerID  bson.ObjectId `json:"ownerID" bson:"_id,omitempty"`
	Location string        `json:"location"`
	Status   string        `json:"status"`
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
}
