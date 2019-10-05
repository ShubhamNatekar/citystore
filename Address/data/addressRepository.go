package data
import(
	"go/pkg.in/mgo.v2"
	"go/pkg.in/mgo.v2/bson"
	"bitbucket.org/shubhamjagdhane/dream_project/src/master/Address/models"
)
type AddressRepository struct{
	C *mgo.Collection
}

func (r *AddressRepository) Create(state *models.StateStruct) error{
	obj_id:=bson.NewObjectId()
	state.StateId=obj_id
	state.State="MAHARASHTRA"
	err:=r.C.Insert(&state)
	return err
}
func (r *AddressRepository) GetRecord() []models.StateStruct{
	var state []models.StateStruct
	iter:=r.C.Find(nil).Iter()
	result:=models.StructState{}
	for iter.Next(&result){
		state =append(state,result)
	}
	return state
}

func (r *AddressRepository) DeleteRecord(id string) error{
	err:=r.C.Remove(bson.M{"_id":bson.ObjectIdHex(id)})
	return err
}

