package services

import (
	"time"

	"github.com/DevEdwinF/smartback.git/internal/config"
)

func CountAttendances() (int64, error) {
	var count int64
	if err := config.DB.Table("attendances").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func CountAttendanceForDay(day time.Time) (int64, error) {
	var count int64
	formattedDay := day.Format("2006-01-02")
	if err := config.DB.Table("attendances").Where("DATE(created_at) = ?", formattedDay).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func CountAttendanceForDayByLate(day time.Time, late bool) (int64, error) {
	var count int64
	formattedDay := day.Format("2006-01-02")
	if err := config.DB.Table("attendances").Where("DATE(created_at) = ? AND late = ?", formattedDay, late).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func CountAttendanceForDayCollaborator(day time.Time, collaboratorId string) (int64, error) {
	var count int64
	formattedDay := day.Format("2006-01-02")
	if err := config.DB.Table("attendances").Where("DATE(created_at) = ? AND fk_collaborator_id = ?", formattedDay, collaboratorId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func TotalCollaboratorsActiveService() (int64, error) {
	var count int64
	if err := config.DB.Table("collaborators").Where("state = 'A'").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
