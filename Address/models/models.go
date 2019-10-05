package models
import (
	"gopkg.in/mgo.v2/bson"
)

type StateStruct struct{
	StateId bson.ObjectId `bson:"_stateid" json:"_stateid,omitempty"`
	State []string `bson: "state" json:"state"`
}
/*
type DistrictStruct struct{
	DistId bson.ObjectId `bson:"_distid" json:"_distid,omitempty"`
	District []string `bson: "district" json:"district"`
=======

type StateStruct struct{
	StateId bson.ObjectId `bson:"_stateid" json:"_stateid,omitempty"`
	State string `bson: "state" json:"state"`
}

type DistrictStruct struct{
	DistId bson.ObjectId `bson:"_distid" json:"_distid,omitempty"`
	District string `bson: "district" json:"district"`
>>>>>>> ec9d03cd3ac5c34279c1ae4f00f8e9c2e9a93382
}

type CityStruct struct{
	CityId bson.ObjectId `bson:"_cityid" json:"_cityid,omitempty"`
<<<<<<< HEAD
	City []string `bson:"city" json:"city"`
=======
	City string `bson:"city" json:"city"`
>>>>>>> ec9d03cd3ac5c34279c1ae4f00f8e9c2e9a93382
}

type PincodeStruct struct{
	PincodeId bson.ObjectId `bson:"_id" json:"_id,omitempty"`
<<<<<<< HEAD
	Pincode []string `bson:"pincode" json:"pincode"`
=======
	Pincode string `bson:"pincode" json:"pincode"`
>>>>>>> ec9d03cd3ac5c34279c1ae4f00f8e9c2e9a93382
}

type PlaceStruct struct{
	PlaceId bson.ObjectId `bson:"_id" json:"_id,omitempty"`
<<<<<<< HEAD
	Place []string `bson:"place" json="place"`
}

	Place string `bson:"place" json="place"`
}
*/
