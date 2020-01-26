package data

import (
	"github.com/aj/my-project/customers/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CustomerRepository struct {
	C *mgo.Collection
}

func (r *CustomerRepository) Create(customer *models.Customer) error {
	obj_id := bson.NewObjectId()
	customer.Id = obj_id
	err := r.C.Insert(&customer)
	return err
}

func (r *CustomerRepository) GetAll() []models.Customer {
	var customers []models.Customer
	iter := r.C.Find(nil).Iter()
	result := models.Customer{}
	for iter.Next(&result) {
		customers = append(customers, result)
	}
	return customers
}

func (r *CustomerRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
