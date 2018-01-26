package repo

import (
	"fmt"

	"github.com/menniti/crm/conf"
	"github.com/menniti/crm/model"
	"gopkg.in/mgo.v2/bson"
)

//GetDealMongo find One deal in MongoDB
func GetDealMongo(dealID string) (deal model.Deal, err error) {
	collection, err := conf.GetMongoCollection("deal")
	if err != nil {
		fmt.Print("[deal] Error to get collection in MONGO", err.Error())
	}
	err = collection.Find(bson.M{"_id": bson.ObjectIdHex(dealID)}).One(&deal)
	if err != nil {
		fmt.Printf("[repo.GetDealMongo] Erro ao retornar _id no MongoDB")
	}
	fmt.Print(&deal)
	return
}

//GetAllDealsMongo find All deals in MongoDB
func GetAllDealsMongo() (results []model.Deal, err error) {
	collection, err := conf.GetMongoCollection("deal")
	if err != nil {
		fmt.Print("[Deal] Error to get All deal collection in MONGO", err.Error())
	}
	err = collection.Find(nil).All(&results)
	if err != nil {
		fmt.Printf("[repo.GetAllDealMongo] Erro ao retornar Todos os deals no MongoDB")
	}
	fmt.Print(&results)
	return
}

//RemoveDealMongo remove one user in MongoDB
func RemoveDealMongo(dealID string) (err error) {
	collection, err := conf.GetMongoCollection("deal")
	if err != nil {
		fmt.Print("[deal] Error to remove deal in MONGO", err.Error())
	}
	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(dealID)})
	if err != nil {
		fmt.Printf("[repo.RemoveDealMongo] Erro ao remover deal no MongoDB")
	}
	return
}

//CreateNewDealMongo create new user in MongoDB
func CreateNewDealMongo(newDeal model.Deal) (err error) {
	collection, err := conf.GetMongoCollection("deal")
	if err != nil {
		fmt.Print("[deal] Error to create new deal in MONGO", err.Error())
	}
	err = collection.Insert(newDeal)
	if err != nil {
		fmt.Printf("[repo.CreateNewDealMongo] Erro ao criar novo deal no MongoDB")
	}
	return
}

//UpdateDealMongo update existent user in MongoDB
func UpdateDealMongo(DealToUpdate model.Deal, dealID string) (err error) {
	collection, err := conf.GetMongoCollection("deal")
	if err != nil {
		fmt.Print("[deal] Error to update deal in MONGO", err.Error())
	}
	queryDeal := bson.M{"_id": bson.ObjectIdHex(dealID)}
	updateDeal := bson.M{"$set": DealToUpdate}
	err = collection.Update(queryDeal, updateDeal)
	if err != nil {
		fmt.Printf("[repo.CreateNewUserMongo] Erro ao atualizar o usuarios no MongoDB")
	}
	return
}
