package controllers
import(
	"gopkg.in/mgo.v2"
	"bitbucket.org/shubhamjagdhane/dream_project/src/master/Address/common"
)

type context struct {
		MongoSession *mgo.Session
}
func (c *Context) Close(){
	c.MongoSession.Close()
}

func (c *Context) DbCollection(name string) *mgo.Collection{
	return c.MongoSession.DB(common.AppConfig.Database)C(name)
}

func NewContext() *Context{
	session:=common.GetSession().Copy()
	context:=&Context{
		MongoSession:session,
	}
	return context
}