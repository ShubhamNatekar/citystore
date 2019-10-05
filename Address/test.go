package main

import(
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/middleware/logger"
)

func main(){
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/test1",testGet1)
	app.Get("/test2",testGet2)

	app.Run(iris.Addr(":8080"))
}

func testGet1(ctx iris.Context){
	ctx.WriteString("This is test-1 string\n")
}

func testGet2(ctx iris.Context){
	ctx.WriteString("This is test-2 string\n")
}
