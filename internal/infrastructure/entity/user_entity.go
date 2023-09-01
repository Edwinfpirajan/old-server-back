package entity

import "time"

type User struct {
	Document string `json:"document"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserData struct {
	User
	FName     string    `json:"f_name"`
	LName     string    `json:"l_name"`
	CreatedAt time.Time `json:"created_at"`
	FkRoleId  int       `json:"rol"`
}

type Users []UserData

type Collaborators struct {
	ID        int       `json:"id_collaborator"`
	Document  string    `json:"document"`
	FName     string    `json:"f_name"`
	LName     string    `json:"l_name"`
	Bmail     string    `json:"bmail"`
	Email     string    `json:"email"`
	Position  string    `json:"position"`
	State     string    `json:"state"`
	Leader    string    `json:"leader"`
	CreatedAt time.Time `json:"date"`
}

type CollaboratorsDataEntity struct {
	Collaborators
	Schedules
}
