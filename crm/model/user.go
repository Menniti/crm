package model

import (
	"gopkg.in/mgo.v2/bson"
)

//User collects data from User
type User struct {
	Name  string        `json:"name" bson:"name"`
	ID    bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Email string        `json:"email" bson:"email"`
	Phone string        `json:"phone" bson:"phone"`
}
