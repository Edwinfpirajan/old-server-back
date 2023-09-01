package controllers

import (
	"net/http"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func GetAllCollaboratorsSchedule(c echo.Context) error {
	collaboratorWithSchedule := []entity.CollaboratorsDataEntity{}

	config.DB.Table("collaborators").Select("*").
		Joins("left join schedule on collaborators.document = schedule.fk_collaborators_document").
		Scan(&collaboratorWithSchedule)

	return c.JSON(http.StatusOK, collaboratorWithSchedule)
}

func AssignSchedulesToCollaborator(c echo.Context) error {
	var schedules []entity.Schedules // Actualizar a la estructura correcta de Schedule

	if err := c.Bind(&schedules); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Formato de datos inválido")
	}

	var collaborator entity.CollaboratorsDataEntity
	err := config.DB.Table("collaborators").Select("*").
		Joins("left join schedules on collaborators.id = schedules.fk_collaborator_id").
		Where(`"collaborators".id = ?`, schedules[0].FkCollaboratorId).
		Order(`"collaborators".id`).
		Limit(1).
		Find(&collaborator).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "No se encuentra el colaborador")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Error en el servidor")
	}

	for _, schedule := range schedules {
		var existingSchedule entity.Schedules
		result := config.DB.Table("schedules").Where("id = ?", schedule.Id).Assign(schedule).FirstOrCreate(&existingSchedule)

		if result.Error != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error al asignar el horario")
		}
	}

	return c.JSON(http.StatusOK, schedules)
}

func DeleteSchedule(c echo.Context) error {
	id := c.Param("id")

	var schedule models.Schedules
	if err := config.DB.First(&schedule, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Schedule not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := config.DB.Delete(&schedule).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "Schedule deleted")
}
