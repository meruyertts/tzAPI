package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"tzAPI/db"
	"tzAPI/models"

	"github.com/gorilla/mux"
)

// GetPeopleHandler handles the HTTP GET request for retrieving a list of people based on query
func GetPeopleHandler(w http.ResponseWriter, r *http.Request) {
	// Extract query parameters from the URL
	name := r.URL.Query().Get("name")
	surname := r.URL.Query().Get("surname")
	patronymic := r.URL.Query().Get("patronymic")
	age := r.URL.Query().Get("age")
	gender := r.URL.Query().Get("gender")
	nationality := r.URL.Query().Get("nationslity")

	// Extract pagination parameters
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	perPage, err := strconv.Atoi(r.URL.Query().Get("perPage"))
	if err != nil || perPage < 1 {
		perPage = 10
	}

	offset := (page - 1) * perPage

	// Build the SQL query based on the provided parameters
	query := "SELECT id, name, surname, patronymic, age, gender, nationality FROM people WHERE TRUE"
	args := []interface{}{}

	if name != "" {
		query += " AND name = $2"
		args = append(args, name)
	}

	if surname != "" {
		query += " AND surname = $3"
		args = append(args, surname)
	}

	if patronymic != "" {
		query += "patronymic = $4"
		args = append(args, patronymic)
	}

	if age != "" {
		query += " AND gender = $5"
		args = append(args, gender)
	}

	if gender != "" {
		query += "gender = $6"
		args = append(args, gender)
	}
	if nationality != "" {
		query += "nationality = $7"
		args = append(args, nationality)
	}
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", perPage, offset)

	// Execute the query and process the results
	rows, err := db.DB.Query(query, args...)
	if err != nil {
		log.Println("Error querying database: ", "error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	people := make([]models.Person, 0)

	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.ID, &person.Name, &person.Surname, &person.Patronymic, &person.Age, &person.Gender, &person.Nationality); err != nil {
			log.Println("Error scanning row: ", "error: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		people = append(people, person)
	}

	// Respond with the retrieved people info
	log.Println("people info was retrieved")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, people)

}

// GetPersonHandler handles the HTTP GET request for retrieving a specific person by ID
func GetPersonHandler(w http.ResponseWriter, r *http.Request) {
	// Extract person ID from the URL parameters
	p1ID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatal(err)
	}

	// Create a Person instance and read the info from the db
	p1 := &models.Person{
		ID: int64(p1ID),
	}
	p2, err := p1.Read()
	if err != nil {
		fmt.Println("error when reading", "error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Respond with the retrieved person info
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(p2)
	if err != nil {
		fmt.Println("error when encoding", "error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("a personal info was retrieved: %+v\n", p2)
	w.WriteHeader(http.StatusOK)
}
