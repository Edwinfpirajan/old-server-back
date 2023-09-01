package controllers

import (
	"errors"
	"net/http"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetAllCollaboratorsController(c echo.Context) error {
	collaborators, err := services.GetAllCollaborators()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "No se encuentra el colaborador"})
	}

	return c.JSON(http.StatusOK, collaborators)
}

func GetAllCollaboratorsHorary(c echo.Context) error {
	collaboratorWithSchedule := []entity.CollaboratorsDataEntity{}

	config.DB.Table("collaborators").Select("*").
		Joins("left join schedule on collaborators.document = schedule.fk_collaborators_document").
		Scan(&collaboratorWithSchedule)

	return c.JSON(http.StatusOK, collaboratorWithSchedule)
}

func GetCollaborator(c echo.Context) error {
	document := c.Param("document")

	collaboratorWithSchedule := []entity.CollaboratorsDataEntity{}

	err := config.DB.Table("collaborators").Select("*").
		Joins("left join schedules on collaborators.id = schedules.fk_collaborator_id").
		Where(`"collaborators".document = ?`, document).
		Order(`"collaborators".document`).
		// Limit(1).
		Find(&collaboratorWithSchedule).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusNotFound, "No se encuentra el colaborador")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Error en el servidor")
	}

	return c.JSON(http.StatusOK, collaboratorWithSchedule)
}

// func SaveCollaborator(c echo.Context) error {
// 	collaborator := entity.CollaboratorsEntity{}

// 	err := c.Bind(&collaborator)

// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}

// 	collaboratorFromDb := models.Collaborators{}

// 	config.DB.Table("collaborators").Where("collaborators.document = ?", collaborator.Document).Scan(&collaborator)

// 	config.DB.Save(&collaboratorFromDb)

// 	return c.JSON(http.StatusCreated, collaborator)
// }

// func CreateOrUpdateSchedule(c echo.Context) error {
// 	document := c.Param("document")

// 	// Obtener los detalles del horario desde el cuerpo de la solicitud
// 	var scheduleData entity.ScheduleEntity
// 	if err := c.Bind(&scheduleData); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Datos del horario inválidos")
// 	}

// 	// Recorrer los días de la semana y crear o actualizar el horario para cada uno
// 	for _, day := range GetDaysOfWeek() {
// 		// Crear una instancia de la estructura de horario
// 		scheduleModel := models.ScheduleModel{
// 			Day:           day,
// 			ArrivalTime:   scheduleData.ArrivalTime,
// 			DepartureTime: scheduleData.DepartureTime,
// 			FkDocument:    scheduleData.FkDocument,
// 		}

// 		// Realizar la creación o actualización del horario en la base de datos
// 		if err := config.DB.Save(&scheduleModel).Error; err != nil {
// 			return echo.NewHTTPError(http.StatusInternalServerError, "Error al guardar el horario")
// 		}
// 	}

// 	return c.JSON(http.StatusOK, "Horarios asignados correctamente")
// }

func DeleteCollaborator(c echo.Context) error {
	id := c.Param("doc")

	employee := models.Collaborators{}

	config.DB.Find(&employee, id)

	if employee.Document != "" {
		config.DB.Delete(employee)
		return c.JSON(http.StatusOK, employee)
	} else {
		return echo.NewHTTPError(http.StatusNotFound, "El colaborador no se encuentra en la base de datos")
	}
}
