package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Address struct {
		Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
		Country     string        `json:"country"`
	}
)
