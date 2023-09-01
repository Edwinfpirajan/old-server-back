package entity

import "time"

// type AttendanceEntity struct {
// 	ID         int64  `json:"document"`
// 	Name       string `json:"name"`
// 	Position   string `json:"position"`
// 	ScheduleId int    `json:"schedule_id"`
// }

type AttendanceEntity struct {
	ID               int64     `json:"id"`
	FkCollaboratorId int       `json:"fk_collaborator_id"`
	Document         string    `json:"document" form:"document"`
	State            string    `json:"state" form:"state"`
	Location         string    `josn:"location" form:"location"`
	Late             *bool     `json:"late"`
	Photo            string    `json:"photo" form:"photo"`
	CreatedAt        time.Time `json:"date"`
}

type UserAttendanceData struct {
	// FkDocumentId int       `json:"document"`
	FkCollaboratorId int       `json:"fk_collaborator_id"`
	Document         string    `json:"document"`
	FName            string    `json:"f_name"`
	LName            string    `json:"l_name"`
	Email            string    `json:"email"`
	Location         string    `json:"location"`
	Arrival          string    `json:"arrival"`
	Departure        string    `json:"departure"`
	Leader           string    `json:"leader"`
	Late             *bool     `json:"late"`
	Photo            string    `json:"photo"`
	CreatedAt        time.Time `json:"date"`
}

type ValidateSchedule struct {
	Id   string `json:"document"`
	Date string `json:"date"`
}

type Translatedcollaborators struct {
	FkCollaboratorId int       `json:"id"`
	Document         string    `json:"document"`
	FName            string    `json:"f_name"`
	LName            string    `json:"l_name"`
	CreatedAt        time.Time `json:"date"`
}
