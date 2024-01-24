package main

import (
	"net/http"
	"tzAPI/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/person", handlers.CreatePersonHandler).Methods(http.MethodPost)
	r.HandleFunc("/people", handlers.GetPeopleHandler).Methods(http.MethodGet)
	r.HandleFunc("/person/{id}", handlers.GetPersonHandler).Methods(http.MethodGet)
	r.HandleFunc("/person/{id}", handlers.UpdatePersonHandler).Methods(http.MethodPut)
	r.HandleFunc("/person/{id}", handlers.DeletePersonHandler).Methods(http.MethodDelete)
	http.ListenAndServe(":8080", r)
}
