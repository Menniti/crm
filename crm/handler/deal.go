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

//GetDeal return deal from database
func GetDeal(ctx *contx.Context) {
	id := ctx.Params(":id")
	deal, err := repo.GetDealMongo(id)
	if err != nil {
		fmt.Printf("[user.GetDealMongo] Erro ao retornar o Deal do MongoDB")
		return
	}
	fmt.Printf("%v\n\r", deal)
	ctx.JSON(http.StatusAccepted, deal)
	return
}

//GetAllDeals return All Deals from database
func GetAllDeals(ctx *contx.Context) {
	deal, err := repo.GetAllDealsMongo()
	if err != nil {
		fmt.Printf("[user.GetAllDealMongo] Erro ao retornar todos os deals do MongoDB")
		return
	}
	fmt.Printf("%v\n\r", deal)
	ctx.JSON(http.StatusAccepted, deal)
	return
}

//UpdateDeal Update deal from database
func UpdateDeal(ctx *contx.Context, req *http.Request) {
	deal := model.Deal{}
	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &deal)
	if err != nil {
		fmt.Printf("[user.NewDeal] Unmarshal nao funcionou")
	}

	id := ctx.Params(":id")
	err = repo.UpdateDealMongo(deal, id)
	if err != nil {
		fmt.Printf("[user.UpdateDeal] Erro ao tentar atualizar o deal")
	}
	fmt.Print("Deal atualizado com sucesso")
}

//NewDeal Add new deal from database
func NewDeal(ctx *contx.Context, req *http.Request) {
	deal := model.Deal{}
	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &deal)
	if err != nil {
		fmt.Printf("[user.NewDeal] Unmarshal nao funcionou")
	}
	err = repo.CreateNewDealMongo(deal)
	if err != nil {
		fmt.Printf("[user.NewDeal] Nao foi possivel adicionar um novo deal no sistema")
	}
}

//DeleteDeal Delete deal from database
func DeleteDeal(ctx *contx.Context, r *http.Request) {
	id := ctx.Params(":id")
	err := repo.RemoveDealMongo(id)
	if err != nil {
		fmt.Printf("[user.DeleteDeal] Erro ao tentar deletar o deal")
	}
	fmt.Print("Deal removido com sucesso")
}
