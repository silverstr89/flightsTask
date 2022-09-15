package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"volume/service"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/flight", service.Flights).Methods("POST")
	http.ListenAndServe(":80", r)
}
