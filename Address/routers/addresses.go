package routers
import(
	"github.com/kataras/iris"
	"bitbuckets.org/shubhamjagdhane/dream_project/src/master/Address/routers"
)
func SetAddressRouters(router *iris.New) *iris.New{
	router.Get("/addresses",controllers.GetAddresses)
	router.Post("/addresses",controllers.InsertAddress)
	return router
}


