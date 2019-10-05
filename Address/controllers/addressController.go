package controller
import(
	"encoding/json"
	"net/http"
	"bitbucket.org/shubhamjagdhane/dream_project/src/master/Address/common"
	"bitbucket.org/shubhamjagdhane/dream_project/src/master/Address/data"
)
func GetAddresses(w http.ResponseWriter,r *http.Request){
	context:=NewContext()
	var dataResource StateResource
	defer context.Close()
	c:=context.DbCollection("state")
	repo:=&data.StructState{c}
	info:=repo.GetAll()
	j,err:=json.Marshal(StateResource{Data:info})
	if err!=nil{
		common.DisplayAppError(w,err,"An unexpeted error has occured",500)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOk)
	w.Write(j)
}
func InsertAddress(w http.ResponseWriter,r *http.Request){
	err:=json.NewDecoder(r.Body).Decode(&dataResource)
	if err!=nil{
		common.DisplayAppError(w,err,"invalid State data",500)
		return
	}
	state:= &dataResource.Data
	context:=NewContext()
	defer context.Close()
	c:=context.DbCollection("state")
	repo:=&data.StructState{c}
	repo.Create(state)
	j,err:=json.Marshal(dataResource)
	if err!=nil{
		common.DisplayAppError(w,err,"An unexpeted error has occured",500)
		return
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOk)
	w.Write(j)
}


