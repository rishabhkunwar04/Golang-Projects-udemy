package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customers{
		{"Rishabh", 25, "Airtel Digital"},
		{"Rahul", 26, "MicroSoft"},
	}
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(customers)

}
func getCustomer_id(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])

}
