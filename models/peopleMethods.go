package models

import (
	"fmt"
	"log"
	"net/http"
	"tzAPI/db"
)

// Create method creates a new person record in the db
func (p *Person) Create() error {
	// Update age, gender, and nationality based on external API
	var err error
	p.Age, err = updateAge(p)
	if err != nil {
		return err
	}
	p.Gender, err = updateGender(p)
	if err != nil {
		return err
	}
	p.Nationality, err = updateNationality(p)
	if err != nil {
		return err
	}
	query := "INSERT INTO people(name, surname, patronymic, age, gender, nationality) VALUES($1, $2, $3, $4, $5, $6) RETURNING id"
	err = db.DB.QueryRow(query, p.Name, p.Surname, p.Patronymic, p.Age, p.Gender, p.Nationality).Scan(&p.ID)
	if err != nil {
		log.Printf("Error executing SQL query: %s\n", query)
		log.Printf("Values: %s, %s, %s, %d, %s, %s\n", p.Name, p.Surname, p.Patronymic, p.Age, p.Gender, p.Nationality)
		log.Println("Scan error:", err)
		err = fmt.Errorf("error message: %s", http.StatusText(http.StatusInternalServerError))
		return err
	}
	return nil
}

// Read retrieves a person's details from the database based on the provided person ID
func (p1 *Person) Read() (*Person, error) {
	query := "SELECT id, name, surname, patronymic, age, gender, nationality FROM people WHERE id = $1"
	row := db.DB.QueryRow(query, p1.ID)

	var p2 Person
	if err := row.Scan(&p2.ID, &p2.Name, &p2.Surname, &p2.Patronymic, &p2.Age, &p2.Gender, &p2.Nationality); err != nil {
		return nil, err
	}
	return &p2, nil

}

func (p1 *Person) Update() error {
	return updateQuery(p1)
}

// Updating a person's information in the db
func updateQuery(p1 *Person) error {
	var query string
	var args []interface{}

	query = "UPDATE people SET "

	if p1.Name != "" {
		query += "name = $2, "
		args = append(args, p1.Name)
	}

	if p1.Surname != "" {
		query += "surname = $3, "
		args = append(args, p1.Surname)
	}

	if p1.Patronymic != "" {
		query += "patronymic = $4, "
		args = append(args, p1.Patronymic)
	}

	if p1.Age != 0 {
		query += "age = $5, "
		args = append(args, p1.Age)
	}
	if p1.Gender != "" {
		query += "gender = $6, "
		args = append(args, p1.Gender)
	}
	if p1.Nationality != "" {
		query += "natinality = $7, "
		args = append(args, p1.Nationality)
	}

	if len(args) > 0 {
		query = query[:len(query)-2]
		query += " WHERE id = $1"
		args = append([]interface{}{p1.ID}, args...)
	} else {
		return nil
	}

	_, err := db.DB.Exec(query, args...)
	return err

}

// deletes the person's information from the database using a named query
func (p *Person) Delete() error {
	_, err := db.DB.NamedExec(`DELETE FROM people WHERE id=:id`,
		map[string]interface{}{
			"id": p.ID,
		})
	return err
}
