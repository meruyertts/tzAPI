package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"tzAPI/models"
)

// CreatePersonHandler handles the HTTP POST request for creating a new person in db
func CreatePersonHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request body to create a new Person struct
	var p1 models.Person
	bb, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	err = json.Unmarshal(bb, &p1)
	if err != nil {
		log.Println("error while unmarshalling ", "error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create a new person using the provided data
	err = p1.Create()
	if err != nil {
		log.Println("error while creating a personal info ", "error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("a personal info was created: %+v\n", p1)
	w.WriteHeader(http.StatusCreated)
}
