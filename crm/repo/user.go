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
		fmt.Print("[user] Error to get All collection in MONGO", err.Error())
	}
	err = collection.Find(nil).All(&results)
	if err != nil {
		fmt.Printf("[repo.GetAllUserMongo] Erro ao retornar Todos os usuarios no MongoDB")
	}
	fmt.Print(&results)
	return
}

//RemoveUserMongo remove one user in MongoDB
func RemoveUserMongo(userID string) (err error) {
	collection, err := conf.GetMongoCollection("user")
	if err != nil {
		fmt.Print("[user] Error to remove user in MONGO", err.Error())
	}
	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(userID)})
	if err != nil {
		fmt.Printf("[repo.RemoveUserMongo] Erro ao remover usuario no MongoDB")
	}
	return
}

//CreateNewUserMongo create new user in MongoDB
func CreateNewUserMongo(newUser model.User) (err error) {
	collection, err := conf.GetMongoCollection("user")
	if err != nil {
		fmt.Print("[user] Error to create new user in MONGO", err.Error())
	}
	err = collection.Insert(newUser)
	if err != nil {
		fmt.Printf("[repo.CreateNewUserMongo] Erro ao criar o usuarios no MongoDB")
	}
	return
}

//UpdateUserMongo update existent user in MongoDB
func UpdateUserMongo(UserToUpdate model.User, userID string) (err error) {
	collection, err := conf.GetMongoCollection("user")
	if err != nil {
		fmt.Print("[user] Error to update user in MONGO", err.Error())
	}
	queryUser := bson.M{"_id": bson.ObjectIdHex(userID)}
	updateUser := bson.M{"$set": UserToUpdate}
	err = collection.Update(queryUser, updateUser)
	if err != nil {
		fmt.Printf("[repo.CreateNewUserMongo] Erro ao atualizar o usuarios no MongoDB")
	}
	return
}
