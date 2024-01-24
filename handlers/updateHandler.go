package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"tzAPI/models"

	"github.com/gorilla/mux"
)

// UpdatePersonHandler handles HTTP requests to update a person's info
func UpdatePersonHandler(w http.ResponseWriter, r *http.Request) {
	// Extract person ID from the URL parameters
	p1ID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}
	p1 := &models.Person{
		ID: int64(p1ID),
	}
	bb, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bb, &p1)
	if err != nil {
		log.Println("error while unmarshalling ", "error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = p1.Update()
	if err != nil {
		fmt.Println("error when updating", "error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("a personal info was updated: %+v\n", p1)
	w.WriteHeader(http.StatusOK)
}
