package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/menniti/crm/lib/contx"
	"github.com/menniti/crm/model"
	"github.com/menniti/crm/repo"
)

//GetCompany return company from database
func GetCompany(ctx *contx.Context) {
	id := ctx.Params(":id")
	company, err := repo.GetCompanyMongo(id)
	if err != nil {
		fmt.Printf("[user.GetCompanyMongo] Erro ao retornar o Company do MongoDB")
		return
	}
	fmt.Printf("%v\n\r", company)
	ctx.JSON(http.StatusAccepted, company)
	return
}

//GetAllCompanies return All companies from database
func GetAllCompanies(ctx *contx.Context) {
	company, err := repo.GetAllCompaniesMongo()
	if err != nil {
		fmt.Printf("[user.GetAllDealMongo] Erro ao retornar todos as companies do MongoDB")
		return
	}
	fmt.Printf("%v\n\r", company)
	ctx.JSON(http.StatusAccepted, company)
	return
}

//UpdateCompany Update company from database
func UpdateCompany(ctx *contx.Context, req *http.Request) {
	company := model.Company{}
	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &company)
	if err != nil {
		fmt.Printf("[user.NewCompany] Unmarshal nao funcionou")
	}

	id := ctx.Params(":id")
	err = repo.UpdateCompanyMongo(company, id)
	if err != nil {
		fmt.Printf("[user.UpdateCompany] Erro ao tentar atualizar o Company")
	}
	fmt.Print("Company atualizado com sucesso")
}

//NewCompany Add company from database
func NewCompany(ctx *contx.Context, req *http.Request) {
	company := model.Company{}
	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &company)
	if err != nil {
		fmt.Printf("[user.NewCompany] Unmarshal nao funcionou")
	}
	err = repo.CreateNewCompanyMongo(company)
	if err != nil {
		fmt.Printf("[user.NewCompany] Nao foi possivel adicionar um novo deal no sistema")
	}
}

//DeleteCompany Delete company from database
func DeleteCompany(ctx *contx.Context) {
	id := ctx.Params(":id")
	err := repo.RemoveCompanyMongo(id)
	if err != nil {
		fmt.Printf("[user.DeleteCompany] Erro ao tentar deletar a Company")
	}
	fmt.Print("Company removido com sucesso")
}
