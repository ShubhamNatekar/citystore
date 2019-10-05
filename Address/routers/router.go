package routers
import(
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)
func InitRoutes() *iris.New{
	router:=iris.New()
	/*router.Logger().SetLevel("debug")
	router.Use(recover.New())
	router.Use(logger.New())*/
	router=SetAddressRouters(router)
	return router
}
