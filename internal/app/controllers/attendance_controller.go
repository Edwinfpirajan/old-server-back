package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AttendanceController struct {
	Service *services.AttendanceService
}

func NewAttendanceController(service *services.AttendanceService) *AttendanceController {
	return &AttendanceController{
		Service: service,
	}
}

func (ac *AttendanceController) SaveRegisterAttendance(c echo.Context) error {
	var attendance entity.AttendanceEntity
	err := c.Bind(&attendance)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = ac.Service.RegisterAttendance(attendance)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Registro de asistencia guardado exitosamente",
	})
}

func (controller *AttendanceController) GetAllAttendance(c echo.Context) error {
	attendance, err := controller.Service.GetAllAttendance()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, attendance)
}
func (controller *AttendanceController) GetAllAttendanceForLate(c echo.Context) error {
	attendance, err := controller.Service.GetAllAttendanceForToLate()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, attendance)
}

func (controller *AttendanceController) GetAttendanceForLeader(c echo.Context) error {
	userToken := c.Get("userToken")

	if userToken == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token de usuario no encontrado")
	}

	token, ok := userToken.(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error al procesar el token")
	}

	claims := token.Claims.(jwt.MapClaims)

	leaderFName, okFName := claims["fName"].(string)
	leaderLName, okLName := claims["lName"].(string)

	if !okFName || !okLName {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": "Este usuario no tiene ningún colaborador asignado",
		})
	}

	leaderFullName := leaderFName + " " + leaderLName

	attendanceService := &services.AttendanceService{}
	attendance, err := attendanceService.GetAttendanceForLeader(leaderFullName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Error obteniendo la asistencia",
		})
	}

	return c.JSON(http.StatusOK, attendance)
}

func (controller *AttendanceController) GetAttendanceForLeaderToLate(c echo.Context) error {
	userToken := c.Get("userToken")

	if userToken == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Token de usuario no encontrado")
	}

	token, ok := userToken.(*jwt.Token)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error al procesar el token")
	}

	claims := token.Claims.(jwt.MapClaims)

	leaderFName, okFName := claims["fName"].(string)
	leaderLName, okLName := claims["lName"].(string)

	if !okFName || !okLName {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"error": "Este usuario no tiene ningún colaborador asignado",
		})
	}

	leaderFullName := leaderFName + " " + leaderLName

	attendanceService := &services.AttendanceService{}
	attendance, err := attendanceService.GetAttendanceForLeaderToLate(leaderFullName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Error obteniendo la asistencia",
		})
	}

	return c.JSON(http.StatusOK, attendance)
}

func ValidateSchedule(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var validateSchedule entity.ValidateSchedule
	err = json.Unmarshal(body, &validateSchedule)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var arrival time.Time

	config.DB.Raw("select arrival from attendance a where fk_document_id = ? and date_format(arrival, '%d-%m-%Y') = date_format(?, '%d-%m-%Y')",
		validateSchedule.Id, validateSchedule.Date).Scan(&arrival)

	return c.JSON(http.StatusOK, arrival)
}

func ValidateCollaboratorController(c echo.Context) error {
	document := c.Param("doc")

	collaborator, err := services.ValidateCollaboratorService(document)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, collaborator)
}

func SaveTranslated(c echo.Context) error {
	var translatedEntity entity.Translatedcollaborators
	err := c.Bind(&translatedEntity)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var collaborator models.Collaborators
	err = config.DB.Model(&collaborator).Where("document = ?", translatedEntity.Document).First(&collaborator).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return echo.NewHTTPError(http.StatusNotFound, "Colaborador no encontrado")
	}

	newTranslatedCollaborator := models.Translatedcollaborators{
		FkCollaboratorId: collaborator.Id,
		CreatedAt:        time.Now(),
	}

	err = config.DB.Create(&newTranslatedCollaborator).Error
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Translado registrado con éxito",
	})
}

func GetAllTranslatedController(c echo.Context) error {
	translatedcollaborators, err := services.GetAllTranslatedService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, translatedcollaborators)
}
