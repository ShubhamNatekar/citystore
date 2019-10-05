package main
import (
	"github.com/kataras/iris"
	//"encoding/json"
	//"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/middleware/logger"
)

func main(){
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	session , err := mgo.Dial("mongodb://127.0.0.1:27017/account")

	if err!=nil{
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic,true)

	c := session.DB("account").C("buyer")

	app.RegisterView(iris.HTML("./view",".html"))

	app.Get("/",func(ctx iris.Context){
		ctx.View("home.html")
	})

	v1 := app.Party("/v1")
		v1.Get("/login/customer",loginPage)
		v1.Get("/signup/customer",signPage)

	app.Get("/allcustomers",func(ctx iris.Context){
		result := []CustInsert{}
		err := c.Find(nil).All(&result)

		if err!=nil{
			panic(err)
		}else{
			ctx.JSON(iris.Map{"Result":result})
		}
	})

//	app.Post("/v1/signup/customer/inserted",func(ctx iris.Context){
	app.Post("/",func(ctx iris.Context){
		if err!=nil{
			ctx.JSON(iris.Map{"response": err.Error()})
		} else{
			doc := CustInsert{
				Id: bson.NewObjectId(),
				MobileNumber: ctx.PostValue("cmobilenumber"),
				Name: ctx.PostValue("cname"),
				Email : ctx.PostValue("cemail"),
				Password: ctx.PostValue("cpass"),
			}
			err=c.Insert(doc)
			if err!=nil{
				panic(err)
			}
			//err = c.Insert(&CustInsert{ctx.PostValue("cmobilenumber"),ctx.PostValue("cname"),ctx.PostValue("cemail")})
			//ctx.WriteString("New User is successfully created")
			ctx.View("home.html")
		}

	})

	app.Run(iris.Addr(":5000"))
}

func loginPage(ctx iris.Context){
	ctx.View("logIn.html")
}
func signPage(ctx iris.Context){
	ctx.View("signUp.html")
}


