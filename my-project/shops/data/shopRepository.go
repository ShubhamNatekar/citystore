package data

import (
	"github.com/aj/my-project/shops/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ShopRepository struct {
	C *mgo.Collection
}

func (r *ShopRepository) Create(shop *models.Shops) error {
	obj_id := bson.NewObjectId()
	shop.Id = obj_id
	err := r.C.Insert(&shop)
	return err
}

func (r *ShopRepository) GetAll() []models.Shops {
	var shops []models.Shops
	iter := r.C.Find(nil).Iter()
	result := models.Shops{}
	for iter.Next(&result) {
		shops = append(shops, result)
	}
	return shops
}

func (r *ShopRepository) Delete(id string) error {
	err := r.C.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
