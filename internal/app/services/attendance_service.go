package services

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"gorm.io/gorm"
)

type AttendanceService struct{}

func NewAttendanceService() *AttendanceService {
	return &AttendanceService{}
}

func (s *AttendanceService) RegisterAttendance(attendance entity.AttendanceEntity) error {
	var collaborator models.Collaborators
	err := config.DB.Model(&collaborator).Where("document = ?", attendance.Document).First(&collaborator).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return errors.New("Colaborador no encontrado")
	}

	fmt.Println(collaborator.Id)

	timeNow := time.Now()

	var schedule models.Schedules
	err = config.DB.Model(&schedule).Where("fk_collaborator_id = ? AND day = ?", collaborator.Id, timeNow.Format("Monday")).First(&schedule).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return errors.New("Horario no encontrado para el colaborador en este día")
	}

	var arrivalScheduled time.Time
	if schedule.ArrivalTime != "" {
		temp, err := time.Parse("15:04:05", schedule.ArrivalTime)
		if err != nil {
			return err
		}
		arrivalScheduled = time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), temp.Hour(), temp.Minute(), temp.Second(), temp.Nanosecond(), timeNow.Location())
	}

	late := false

	if !arrivalScheduled.IsZero() && timeNow.After(arrivalScheduled.Add(5*time.Minute)) {
		late = true
	}

	var validateAttendance models.Attendance
	err = config.DB.Model(&validateAttendance).
		Where("fk_collaborator_id = ? AND date(created_at) = ?", collaborator.Id, timeNow.Format("2006-01-02")).
		First(&validateAttendance).Error

	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	folderPath := "attendance_photos"
	err = os.MkdirAll(folderPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	imagenCodificadaEnBase64 := attendance.Photo

	marca := ";base64,"
	indice := strings.Index(imagenCodificadaEnBase64, marca)
	if indice != -1 {
		imagenCodificadaEnBase64 = imagenCodificadaEnBase64[indice+len(marca):]
	}

	decodificado, err := base64.StdEncoding.DecodeString(imagenCodificadaEnBase64)
	if err != nil {
		log.Fatal(err)
	}

	imagen, _, err := image.Decode(bytes.NewReader(decodificado))
	if err != nil {
		log.Fatal(err)
	}

	photoName := fmt.Sprintf("%s%d.png", "1150856537", time.Now().Unix())

	archivo, err := os.Create(fmt.Sprintf("%s/%s", folderPath, photoName))
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(archivo, imagen)
	if err != nil {
		log.Fatal(err)
	}
	err = archivo.Close()
	if err != nil {
		log.Fatal(err)
	}

	attendance.Photo = photoName

	switch attendance.State {
	case "arrival":
		if validateAttendance.ID != 0 || validateAttendance.Arrival.Valid {
			return errors.New("Ya se ha registrado la entrada")
		}

		modelsAttendance := models.Attendance{
			FkCollaboratorID: collaborator.Id,
			Photo:            attendance.Photo,
			Location:         attendance.Location,
			Arrival:          sql.NullString{String: timeNow.Format("15:04:05"), Valid: true},
			Late:             late,
			CreatedAt:        timeNow,
		}
		err = config.DB.Create(&modelsAttendance).Error
		if err != nil {
			return err
		}

		return nil

	case "departure":
		if validateAttendance.ID == 0 {
			return errors.New("Debe registrar la entrada primero")
		}
		if validateAttendance.Departure.Valid {
			return errors.New("Ya se ha registrado la salida")
		}

		validateAttendance.Departure = sql.NullString{String: timeNow.Format("15:04:05"), Valid: true}

		err = config.DB.Updates(&validateAttendance).Error
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("Estado inválido")
}

func (service *AttendanceService) GetAllAttendance() ([]entity.UserAttendanceData, error) {
	attendance := []entity.UserAttendanceData{}
	err := config.DB.Table("attendances a").
		Select("c.f_name, c.l_name, c.email, c.document, a.*").
		Joins("INNER JOIN collaborators c on c.id = a.fk_collaborator_id").
		Find(&attendance).Error
	if err != nil {
		return nil, err
	}

	folderPath := "attendance_photos"

	for i := range attendance {
		photoName := attendance[i].Photo
		imagePath := filepath.Join(folderPath, photoName)

		imageData, err := ioutil.ReadFile(imagePath)
		if err != nil {
			return nil, err
		}

		base64Image := base64.StdEncoding.EncodeToString(imageData)

		attendance[i].Photo = base64Image
	}

	return attendance, nil
}

func (service *AttendanceService) GetAttendanceForLeader(leaderFullName string) ([]entity.UserAttendanceData, error) {
	attendance := []entity.UserAttendanceData{}
	err := config.DB.Table("attendances a").
		Select("c.f_name, c.l_name, c.email, c.leader, c.document, a.*").
		Joins("INNER JOIN collaborators c ON c.id = a.fk_collaborator_id").
		Joins("INNER JOIN users u ON CONCAT(u.f_name, ' ', u.l_name) = c.leader").
		Where("c.leader = ?", leaderFullName).
		Find(&attendance).Error
	if err != nil {
		return nil, err
	}

	folderPath := "attendance_photos"

	for i := range attendance {
		photoName := attendance[i].Photo
		imagePath := filepath.Join(folderPath, photoName)

		imageData, err := ioutil.ReadFile(imagePath)
		if err != nil {
			return nil, err
		}

		base64Image := base64.StdEncoding.EncodeToString(imageData)

		attendance[i].Photo = base64Image
	}

	return attendance, nil
}

func (service *AttendanceService) GetAllAttendanceForToLate() ([]entity.UserAttendanceData, error) {
	attendance := []entity.UserAttendanceData{}
	err := config.DB.Table("collaborators c").
		Select("c.f_name, c.l_name, c.email, c.leader, c.document, a.*").
		Joins("INNER JOIN users u ON CONCAT(u.f_name, ' ', u.l_name) = c.leader").
		Joins("INNER JOIN attendances a ON c.id = a.fk_collaborator_id").
		Where("EXISTS (SELECT 1 FROM attendances WHERE fk_collaborator_id = c.id AND late = TRUE HAVING COUNT(*) > 2)").
		Find(&attendance).Error
	if err != nil {
		return nil, err
	}

	folderPath := "attendance_photos"

	for i := range attendance {
		photoName := attendance[i].Photo
		imagePath := filepath.Join(folderPath, photoName)

		imageData, err := ioutil.ReadFile(imagePath)
		if err != nil {
			return nil, err
		}

		base64Image := base64.StdEncoding.EncodeToString(imageData)

		attendance[i].Photo = base64Image
	}

	return attendance, nil
}

func (service *AttendanceService) GetAttendanceForLeaderToLate(leaderFullName string) ([]entity.UserAttendanceData, error) {
	attendance := []entity.UserAttendanceData{}
	err := config.DB.Table("collaborators c").
		Select("c.f_name, c.l_name, c.email, c.leader, c.document, a.*").
		Joins("INNER JOIN users u ON CONCAT(u.f_name, ' ', u.l_name) = c.leader").
		Joins("INNER JOIN attendances a ON c.id = a.fk_collaborator_id").
		Where("c.leader = ?", leaderFullName).
		Where("EXISTS (SELECT 1 FROM attendances WHERE fk_collaborator_id = c.id AND late = TRUE HAVING COUNT(*) > 2)").
		Find(&attendance).Error
	if err != nil {
		return nil, err
	}

	folderPath := "attendance_photos"

	for i := range attendance {
		photoName := attendance[i].Photo
		imagePath := filepath.Join(folderPath, photoName)

		imageData, err := ioutil.ReadFile(imagePath)
		if err != nil {
			return nil, err
		}

		base64Image := base64.StdEncoding.EncodeToString(imageData)

		attendance[i].Photo = base64Image
	}

	return attendance, nil
}

// func (service *AttendanceService) GetAllAttendanceTest() ([]entity.UserAttendanceData, error) {
// 	attendance := []entity.UserAttendanceData{}
// 	err := config.DB.Table("attendances a").
// 		Select("c.f_name, c.l_name, c.email, c.document, a.*").
// 		Joins("INNER JOIN collaborators c on c.id = a.fk_collaborator_id").
// 		Find(&attendance).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	folderPath := "attendance_photos"

// 	for i := range attendance {
// 		photoName := attendance[i].Photo
// 		imagePath := filepath.Join(folderPath, photoName)

// 		imageData, err := ioutil.ReadFile(imagePath)
// 		if err != nil {
// 			return nil, err
// 		}

// 		base64Image := base64.StdEncoding.EncodeToString(imageData)

// 		attendance[i].Photo = base64Image
// 	}

// 	return attendance, nil
// }
