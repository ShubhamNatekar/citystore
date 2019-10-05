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

	app.Get("/",test)

	app.Run(iris.Addr(":8080"))
}

func test(ctx iris.Context){
	ctx.WriteString("This is test string\n")
}

