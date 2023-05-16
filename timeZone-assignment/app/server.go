package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	mux := mux.NewRouter()
	mux.HandleFunc("/api/time", CurrentTime)
	mux.HandleFunc("/api/time/{para_m}", regionalTime)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}
