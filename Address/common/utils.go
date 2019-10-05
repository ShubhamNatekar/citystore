package common
import(
	"encoding/json"
	"net/http"
	"log"
	"os"
	"time"
	"gopkg.in/mgo.v2"
)
type(
	appError struct {
			Error	string `json:"error"`
			Message	string `json:"message"`
			HttpStatus	int `json:"status"`
	}
	errorResource struct{
		Data appError `json:"data"`
	}
	configuration struct{
		Server,MongoDBHost,DBUser,Database string
	}
)

func DisplayAppError(w http.ResponseWriter,handlerError error,message string,code int){
	errObj:=appError{
			Error :	handlerError.Error(),
			Message :	message,
			HttpStatus :	code,
	}
	log.Printf("appError: %s\n",handlerError)
	w.Header().Set("Content-Type","application/json; charset=utf-8")
	w.WriteHeader(code)
	if j,err:=json.Marshal(errorResource{Data:errObj}); err==nil{
		w.Write(j)
	}
}
var AppConfig configuration

func initConfig(){
	file,err:=os.Open("common/config.json")
	defer file.Close()
	if err!=nil{
		log.Fatalf("[loadConfig]: %s\n",err)
	}
	decoder:=json.NewDecoder(file)
	AppConfig=configuration{}
	err=decoder.Decoder(&AppConfig)
	if err!=nil{
		log.Fatalf("[loadConfig]: %s\n",err)
	}
	loadConfigFromEnvironment(&AppConfig)
}

func loadConfigFromEnvironment(appConfig *configuration){
	server,ok:=os.LookupEnv("AllAddresses")
	if ok{
		appConfig.Server=server
		log.Printf("[INFO]:Server Information loaded from Environment ")
	}
	mongoDBHost,ok:=os.LookupEnv("MongoDBHost")
	if ok{
		appConfig.MongoDBHost=mongoDBHost
		log.Printf("[INFO]:Server Information loaded from Environment ")
	}
	mongoDBUser,ok:=os.LookupEnv("MongoDBUser")
	if ok{
		appConfig.MongoDBUser=mongoDBUser
		log.Printf("[INFO]:Server Information loaded from Environment ")
	}
	mongoDBPwd,ok:=os.LookupEnv("MongoDBPwd")
	if ok{
		appConfig.MongoDBPwd=mongoDBPwd
		log.Printf("[INFO]:Server Information loaded from Environment ")
	}
	database,ok:=os.LookupEnv("Database")
	if ok{
		appConfig.Database=database
		log.Printf("[INFO]:Server Information loaded from Environment ")
	}
}
var session *mgo.Session

func getSession() *mgo.Session{
	if session ==nil {
		var err error
		session,err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:	[]string{AppConfig.MongoDBHost},
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout: 60*time.Second,
		})
		if err!=nil{
				log.Fatalf("[GetSession]: %s\n",err)
		}
	}
	return session
}
func createDbSession(){
	session,err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:	[]string{AppConfig.MongoDBHost},
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout: 60*time.Second,
	})
	if err!=nil{
		log.Fatalf("[createDbSession]: %s\n",err)
	}
}
