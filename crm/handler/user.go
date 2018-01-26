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

//GetUser return user from database
func GetUser(ctx *contx.Context) {

	fmt.Printf("%v\n\r", ctx.Params(":id"))
	id := ctx.Params(":id")
	user, err := repo.GetUserMongo(id)
	if err != nil {
		fmt.Printf("[user.GetUserMongo] Erro ao retornar o user do MongoDB")
		return
	}
	fmt.Printf("%v\n\r", user)
	ctx.JSON(http.StatusAccepted, user)
	return
}

//GetAllUsers return All Users from database
func GetAllUsers(ctx *contx.Context) {

	user, err := repo.GetAllUserMongo()
	if err != nil {
		fmt.Printf("[user.GetUserMongo] Erro ao retornar o user do MongoDB")
		return
	}
	fmt.Printf("%v\n\r", user)
	ctx.JSON(http.StatusAccepted, user)
	return
}

//NewUser Add new user from database
func NewUser(ctx *contx.Context, req *http.Request) {

	user := model.User{}
	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Printf("[user.NewUser] Unmarshal nao funcionou")
	}
	err = repo.CreateNewUserMongo(user)
	if err != nil {
		fmt.Printf("[user.NewUser] Nao foi possivel adicionar um novo usuario no sistema")
	}
}

//DeleteUser Delete user from database
func DeleteUser(ctx *contx.Context) {
	id := ctx.Params(":id")
	err := repo.RemoveUserMongo(id)
	if err != nil {
		fmt.Printf("[user.DeleteUser] Erro ao tentar deletar o usuario")
	}
	fmt.Print("Usuario removido com sucesso")
}

//UpdateUser Update user from database
func UpdateUser(ctx *contx.Context, req *http.Request) {

	user := model.User{}
	body, err := ioutil.ReadAll(req.Body)
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Printf("[user.NewUser] Unmarshal nao funcionou")
	}

	id := ctx.Params(":id")
	err = repo.UpdateUserMongo(user, id)
	if err != nil {
		fmt.Printf("[user.UpdateUser] Erro ao tentar atualizar o usuario")
	}
	fmt.Print("Usuario atualizado com sucesso")
}
