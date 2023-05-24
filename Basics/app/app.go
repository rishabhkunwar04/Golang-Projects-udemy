package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customers struct {
	Name    string `json:"Full Name"`
	Age     int
	Company string
}

func Start() {
	//mux := http.NewServeMux() //degfault multiplexer
	mux := mux.NewRouter() //using gorilla mux multiplexer
	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Go lang Application")
	})
	mux.HandleFunc("/getCustomers", getAllCustomers)
	mux.HandleFunc("/customer/{customer_id}", getCustomer_id)

	//starting server
	http.ListenAndServe("localhost:8000", mux)

	/*
		 //Default multiplexer
		//response writer help us sendind info back to client  and request accept the url parameters
		http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Go lang Application")
		})
		http.HandleFunc("/getCustomers", getAllCustomers)

		//starting server
		http.ListenAndServe("localhost:8000", nil) //in this case handler is nil so default multiplexer will be called

	*/
}
