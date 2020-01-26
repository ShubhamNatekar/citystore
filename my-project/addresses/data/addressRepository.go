package data

import (
	"github.com/aj/my-project/addresses/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AddressRepository struct {
	C *mgo.Collection
}

func (r *AddressRepository) Create(address *models.Address) error {
	obj_id := bson.NewObjectId()
	address.Id = obj_id
	err := r.C.Insert(&address)
	return err
}

func (r *AddressRepository) GetAll() []models.Address {
	var address []models.Address
	iter := r.C.Find(nil).Iter()
	result := models.Address{}
	for iter.Next(&result) {
		address = append(address, result)
	}
	return address
}

func (r *AddressRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
