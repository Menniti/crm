package handler

import (
	"fmt"
	"net/http"
)

//GetCompany return company from database
func GetCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetCompany")
}

//GetAllCompanies return All companies from database
func GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GetAllCompanies")
}

//UpdateCompany Update company from database
func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "UpdateCompany")
}

//NewCompany Add company from database
func NewCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "NewCompany")
}

//DeleteCompany Delete company from database
func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "DeleteCompany")
}
