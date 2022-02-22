package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/adamatti/graphql-go/db"
	"github.com/adamatti/graphql-go/graphql"
	"github.com/adamatti/graphql-go/graphql2"

	mux "github.com/gorilla/mux"
)

func getCountriesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Get countries called")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.GetCountryList())
}

func getStatesHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Get states called")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.GetStateList())
}

func buildRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/countries", getCountriesHandler).Methods("GET")
	router.HandleFunc("/states", getStatesHandler).Methods("GET")
	router.Handle("/graphql", graphql.GetHandle())
	router.Handle("/graphql2", graphql2.GetHandle())
	return router
}

func main() {
	log.Println("App starting")
	go http.ListenAndServe(":8000", buildRouter())

	select{}
}