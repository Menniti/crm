package repo

import (
	"fmt"

	"github.com/menniti/crm/conf"
	"github.com/menniti/crm/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MongoSection collects section mongodb
var MongoSection *mgo.Session

//GetUserMongo find One user in MongoDB
func GetUserMongo(userID string) (user model.User, err error) {
	collection, err := conf.GetMongoCollection("user")
	if err != nil {
		fmt.Print("[user] Error to get collection in MONGO", err.Error())
	}
	err = collection.Find(bson.M{"_id": bson.ObjectIdHex(userID)}).One(&user)
	if err != nil {
		fmt.Printf("[repo.GetUserMongo] Erro ao retornar _id no MongoDB")
	}
	fmt.Print(&user)
	return
}

//GetAllUserMongo find All users in MongoDB
func GetAllUserMongo() (results []model.User, err error) {
	collection, err := conf.GetMongoCollection("user")
	if err != nil {
		fmt.Print("[user] Error to get collection in MONGO", err.Error())
	}
	err = collection.Find(nil).All(&results)
	if err != nil {
		fmt.Printf("[repo.GetAllUserMongo] Erro ao retornar Todos os usuarios no MongoDB")
	}
	fmt.Print(&results)
	return
}
