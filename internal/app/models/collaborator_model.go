package models

import "time"

type Collaborators struct {
	Id        int
	Document  string
	FName     string
	LName     string
	Email     string
	Bmail     string
	State     string
	Position  string
	CreatedAt time.Time
}

type CollaboratorsData struct {
	Collaborators
	Schedules
}

// type Employe struct {
// 	ID         int       `json:"id" gorm:"primary_key;auto_increment"`
// 	PinEmploye string    `json:"pinEmploye" gorm:"FOREIGNKEY:PinEmploye" `
// 	FirstName  string    `json:"first_name" `
// 	LastName   string    `json:"last_name"`
// 	Company    string    `json:"company"`
// 	Position   string    `json:"position"`
// 	ScheduleId int       `json:"schedule_id"`
// 	CreatedAt  time.Time `json:"fechacreacion"`
// }
