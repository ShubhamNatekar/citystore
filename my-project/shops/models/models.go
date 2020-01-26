package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
	Shops struct {
		Id       bson.ObjectId `bson:"_id,omitempty" json:"id"`
		ShopName     string        `json:"shopname"`
	}
)
