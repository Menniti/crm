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

//GetClient return client from database
func GetClient(ctx *contx.Context) {
	id := ctx.Params(":id")
	client, err := repo.GetClientMongo(id)
	if err != nil {
		fmt.Printf("[user.GetClientMongo] Erro ao retornar o Client do MongoDB")
		return
	}
	fmt.Printf("%v\n\r", client)
	ctx.JSON(http.StatusAccepted, client)
	return
}

//GetAllClients return All clients from database
func GetAllClients(ctx *contx.Context) {
	client, err := repo.GetAllClientsMongo()
	if err != nil {
		fmt.Printf("[user.GetAllClientsMongo] Erro ao retornar todos clients do MongoDB")
		return
	}
	fmt.Printf("%v\n\r", client)
	ctx.JSON(http.StatusAccepted, client)
	return
}

//UpdateClient Update client from database
func UpdateClient(ctx *contx.Context, req *http.Request) {
	client := model.Client{}
	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &client)
	if err != nil {
		fmt.Printf("[user.NewClient] Unmarshal nao funcionou")
	}

	id := ctx.Params(":id")
	err = repo.UpdateClientMongo(client, id)
	if err != nil {
		fmt.Printf("[user.UpdateClient] Erro ao tentar atualizar o Client")
	}
	fmt.Print("Client atualizado com sucesso")
}

//NewClient Add new client from database
func NewClient(ctx *contx.Context, req *http.Request) {
	client := model.Client{}
	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &client)
	if err != nil {
		fmt.Printf("[user.NewClient] Unmarshal nao funcionou")
	}
	err = repo.CreateNewClientMongo(client)
	if err != nil {
		fmt.Printf("[user.NewClient] Nao foi possivel adicionar um novo client no sistema")
	}
}

//DeleteClient Delete client from database
func DeleteClient(ctx *contx.Context) {
	id := ctx.Params(":id")
	err := repo.RemoveClientMongo(id)
	if err != nil {
		fmt.Printf("[user.DeleteClient] Erro ao tentar deletar o Client")
	}
	fmt.Print("Client removido com sucesso")
}
