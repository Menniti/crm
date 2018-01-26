package model

import "gopkg.in/mgo.v2/bson"

//Deal collects data from Deal
type Deal struct {
	Name      string        `json:"name" bson:"name"`
	DealOwner string        `json:"dealOwner" bson:"dealOwner"`
	Status    string        `json:"status" bson:"status"`
	Value     int           `json:"value" bson:"value"`
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	OwnerID   bson.ObjectId `json:"ownerID" bson:"ownerID,omitempty"`
}
