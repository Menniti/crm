package handler

import (
	"fmt"
	"net/http"

	"github.com/menniti/crm/lib/contx"
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
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetAllUsers")
}

//UpdateUser Update user from database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UpdateUser")
}

//NewUser Add new user from database
func NewUser(ctx *contx.Context) {

}

//DeleteUser Delete user from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "DeleteUser")
}
