package models

type Headquarter struct {
	Id   int    `json:"id" param:"id"`
	Name string `json:"hearquarter"`
}

type Headquarters []Headquarter
