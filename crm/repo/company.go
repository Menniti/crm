package repo

import (
	"fmt"

	"github.com/menniti/crm/conf"
	"github.com/menniti/crm/model"
	"gopkg.in/mgo.v2/bson"
)

//GetCompanyMongo find One company in MongoDB
func GetCompanyMongo(companyID string) (company model.Company, err error) {
	collection, err := conf.GetMongoCollection("company")
	if err != nil {
		fmt.Print("[company] Error to get collection in MONGO", err.Error())
	}
	err = collection.Find(bson.M{"_id": bson.ObjectIdHex(companyID)}).One(&company)
	if err != nil {
		fmt.Printf("[repo.GetCompanyMongo] Erro ao retornar _id no MongoDB")
	}
	fmt.Print(&company)
	return
}

//GetAllCompaniesMongo find All Companies in MongoDB
func GetAllCompaniesMongo() (results []model.Company, err error) {
	collection, err := conf.GetMongoCollection("company")
	if err != nil {
		fmt.Print("[company] Error to get All Companies collection in MONGO", err.Error())
	}
	err = collection.Find(nil).All(&results)
	if err != nil {
		fmt.Printf("[repo.GetAllCompanyMongo] Erro ao retornar Todos as companies no MongoDB")
	}
	fmt.Print(&results)
	return
}

//RemoveCompanyMongo remove one company in MongoDB
func RemoveCompanyMongo(companyID string) (err error) {
	collection, err := conf.GetMongoCollection("company")
	if err != nil {
		fmt.Print("[company] Error to remove company in MONGO", err.Error())
	}
	err = collection.Remove(bson.M{"_id": bson.ObjectIdHex(companyID)})
	if err != nil {
		fmt.Printf("[repo.RemoveCompanyMongo] Erro ao remover company no MongoDB")
	}
	return
}

//CreateNewCompanyMongo create new company in MongoDB
func CreateNewCompanyMongo(newCompany model.Company) (err error) {
	collection, err := conf.GetMongoCollection("company")
	if err != nil {
		fmt.Print("[deal] Error to create new deal in MONGO", err.Error())
	}
	err = collection.Insert(newCompany)
	if err != nil {
		fmt.Printf("[repo.CreateNewCompanyMongo] Erro ao criar nova company no MongoDB")
	}
	return
}

//UpdateCompanyMongo update existent user in MongoDB
func UpdateCompanyMongo(CompanyToUpdate model.Company, companyID string) (err error) {
	collection, err := conf.GetMongoCollection("company")
	if err != nil {
		fmt.Print("[company] Error to update company in MONGO", err.Error())
	}
	queryCompany := bson.M{"_id": bson.ObjectIdHex(companyID)}
	updateCompany := bson.M{"$set": CompanyToUpdate}
	err = collection.Update(queryCompany, updateCompany)
	if err != nil {
		fmt.Printf("[repo.CreateNewUserMongo] Erro ao atualizar o company no MongoDB")
	}
	return
}
