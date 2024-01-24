package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func updateAge(p *Person) (int, error) {
	var newAgify AgifyStruct
	resp, err := http.Get("https://api.agify.io/?name=" + p.Name)
	if err != nil {
		fmt.Println(http.StatusNotFound)
		return 0, err
	}
	defer resp.Body.Close()
	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	err = json.Unmarshal(bb, &newAgify)
	if err != nil {
		return 0, err
	}
	return newAgify.Age, nil
}

func updateGender(p *Person) (string, error) {
	var newGenderize GenderizeStruct
	resp, err := http.Get("https://api.genderize.io/?name=" + p.Name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(bb, &newGenderize)
	if err != nil {
		return "", err
	}
	return newGenderize.Gender, nil
}

func updateNationality(p *Person) (string, error) {
	var newNationalize Nationalize
	resp, err := http.Get("https://api.nationalize.io/?name=" + p.Name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bb, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(bb, &newNationalize)
	if err != nil {
		return "", err
	}
	return countryID(newNationalize.Country), nil
}

func countryID(countries []Countrify) string {
	max, countryMax := countries[0].Probability, countries[0].Country_ID
	for i := 1; i < len(countries); i++ {
		if countries[i].Probability > max {
			max = countries[i].Probability
			countryMax = countries[i].Country_ID
		}
	}
	return countryMax
}
