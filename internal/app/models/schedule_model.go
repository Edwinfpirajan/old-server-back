package models

type Schedules struct {
	Id               int `gorm:"primaryKey"`
	Day              string
	ArrivalTime      string
	DepartureTime    string
	FkCollaboratorId int `gorm:"column:fk_collaborator_id"`
}
