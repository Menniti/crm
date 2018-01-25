package handler

import (
	"fmt"
	"net/http"
)

//GetDeal return deal from database
func GetDeal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetDeal")
}

//GetAllDeals return All Deals from database
func GetAllDeals(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetAllDeals")
}

//UpdateDeal Update deal from database
func UpdateDeal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UpdateDeal")
}

//NewDeal Add new deal from database
func NewDeal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NewDeal")
}

//DeleteDeal Delete deal from database
func DeleteDeal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "DeleteDeal")
}
