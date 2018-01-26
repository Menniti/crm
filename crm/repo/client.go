package repo

import (
	"fmt"

	"github.com/menniti/crm/conf"
	"github.com/menniti/crm/model"
	"gopkg.in/mgo.v2/bson"
)

//GetClientMongo find One Client in MongoDB
func GetClientMongo(clientID string) (client model.Client, err error) {
	collection, err := conf.GetMongoCollection("client")
	if err != nil {
		fmt.Print("[company] Error to get collection in MONGO", err.Error())
	}
	err = collection.Find(bson.M{"_id": bson.ObjectIdHex(clientID)}).One(&client)
	if err != nil {
		fmt.Printf("[repo.GetClientMongo] Erro ao retornar _id no MongoDB")
	}
	fmt.Print(&client)
	return
}

//GetAllClientsMongo find All Clients in MongoDB
func GetAllClientsMongo() (results []model.Client, err error) {
	collection, err := conf.GetMongoCollection("client")
	if err != nil {
		fmt.Print("[company] Error to get All Client collection in MONGO", err.Error())
	}
	err = collection.Find(nil).All(&results)
	if err != nil {
		fmt.Printf("[repo.GetAllClientsMongo] Erro ao retornar Todos os clients no MongoDB")
	}
	fmt.Print(&results)
	return
}

//RemoveClientMongo remove one client in MongoDB
func RemoveClientMongo(clientID string) (err error) {
	collection, err := conf.GetMongoCollection("client")
	if err != nil {
		fmt.Print("[company] Error to remove client in MONGO", err.Error())
	}
	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(clientID)})
	if err != nil {
		fmt.Printf("[repo.RemoveClientMongo] Erro ao remover Client no MongoDB")
	}
	return
}

//CreateNewClientMongo create new client in MongoDB
func CreateNewClientMongo(newClient model.Client) (err error) {
	collection, err := conf.GetMongoCollection("client")
	if err != nil {
		fmt.Print("[deal] Error to create new client in MONGO", err.Error())
	}
	err = collection.Insert(newClient)
	if err != nil {
		fmt.Printf("[repo.CreateNewClientMongo] Erro ao criar novo client no MongoDB")
	}
	return
}

//UpdateClientMongo update existent client in MongoDB
func UpdateClientMongo(ClientToUpdate model.Client, companyID string) (err error) {
	collection, err := conf.GetMongoCollection("client")
	if err != nil {
		fmt.Print("[client] Error to update client in MONGO", err.Error())
	}
	queryClient := bson.M{"_id": bson.ObjectIdHex(companyID)}
	updateClient := bson.M{"$set": ClientToUpdate}
	err = collection.Update(queryClient, updateClient)
	if err != nil {
		fmt.Printf("[repo.CreateNewClientMongo] Erro ao atualizar o client no MongoDB")
	}
	return
}
