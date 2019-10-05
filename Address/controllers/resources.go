package controllers
import(
	"bitbucket.org/shubhamjagdhane/dream_project/src/master/Address/models"
)
type (
	AddressResource struct{
		Data []models.StateStruct `json:"data"`
	}
	AddressResource struct{
		Data models.StateStruct `json:"data"`
	}

)
