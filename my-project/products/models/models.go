package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	ProductTypes struct {
		Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ProductType     string        `json:"ptype"`
	}
)
