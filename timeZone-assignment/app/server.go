package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func Start() {
	mux := mux.NewRouter()
	mux.HandleFunc("/api/time", getTime)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))

}

func getTime(w http.ResponseWriter, r *http.Request) {

	response := make(map[string]string, 0)
	tz := r.URL.Query().Get("tz")
	timezones := strings.Split(tz, ",")

	if len(timezones) <= 1 {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("invalid timezone %s", loc)))
		} else {
			response[loc.String()] = time.Now().In(loc).String()
			w.Header().Add("Content-type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	} else {
		for _, val := range timezones {
			loc, err := time.LoadLocation(val)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("invalid timezones %s", val)))
			} else {
				response[val] = time.Now().In(loc).String()
				w.Header().Add("Content-type", "application/json")
				json.NewEncoder(w).Encode(response)
			}
		}

	}

}
