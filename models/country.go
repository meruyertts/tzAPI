package models

type Countrify struct {
	Country_ID  string  `json:"country_id"`
	Probability float64 `json:"probability"`
}
type Nationalize struct {
	Count   int         `json:"count"`
	Name    string      `json:"name"`
	Country []Countrify `json:"country"`
}
