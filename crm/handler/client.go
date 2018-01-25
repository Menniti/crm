package handler

import (
	"fmt"
	"net/http"
)

//GetClient return client from database
func GetClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetClient")
}

//GetAllClients return All clients from database
func GetAllClients(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetAllClients")
}

//UpdateClient Update client from database
func UpdateClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UpdateClient")
}

//NewClient Add new client from database
func NewClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NewClient")
}

//DeleteClient Delete client from database
func DeleteClient(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "DeleteClient")
}
