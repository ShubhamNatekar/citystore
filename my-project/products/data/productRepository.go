package data

import (
	"github.com/aj/my-project/products/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductRepository struct {
	C *mgo.Collection
}

func (r *ProductRepository) Create(product *models.ProductTypes) error {
	obj_id := bson.NewObjectId()
	product.Id = obj_id
	err := r.C.Insert(&product)
	return err
}

func (r *ProductRepository) GetAll() []models.ProductTypes {
	var products []models.ProductTypes
	iter := r.C.Find(nil).Iter()
	result := models.ProductTypes{}
	for iter.Next(&result) {
		products = append(products, result)
	}
	return products
}

func (r *ProductRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
