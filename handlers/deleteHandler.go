package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tzAPI/models"

	"github.com/gorilla/mux"
)

// DeletePersonHandler handles HTTP requests to delete a person's info
func DeletePersonHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the person's ID from the request path parameters
	p1ID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println("error on deleting a personal info ", "error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	p1 := models.Person{
		ID: int64(p1ID),
	}
	err = p1.Delete()
	if err != nil {
		fmt.Println("error when deleting", "error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("a personal info was deleted")
	w.WriteHeader(http.StatusNoContent)

}
