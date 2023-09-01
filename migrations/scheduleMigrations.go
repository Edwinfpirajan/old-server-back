package migrations

import (
	"fmt"
	"net/http"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

type ScheduleMigrationService struct{}

func NewScheduleMigrationService() *ScheduleMigrationService {
	return &ScheduleMigrationService{}
}

func (s *ScheduleMigrationService) MigrationSchedule(schedule models.Schedules) {
}

// func UploadExcelAndAssignSchedules(c echo.Context) error {
// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "Error al cargar el archivo")
// 	}

// 	excelFile, err := excelize.OpenFile(file.Filename)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "Error al abrir el archivo Excel")
// 	}

// 	rows, err := excelFile.GetRows("horario")
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "Error al leer las filas del archivo Excel")
// 	}

// 	for _, row := range rows {
// 		document := row[0]
// 		day := row[1]
// 		arrivalTime := row[2]
// 		departureTime := row[3]

// 		query := fmt.Sprintf(`INSERT INTO "schedules" ("day","arrival_time","departure_time","fk_collaborator_id")
//                              SELECT ?, ?, ?, c."id"
//                              FROM "collaborators" c
//                              WHERE c."document" = ?
//                              RETURNING "id"`)
// 		var scheduleID int
// 		err := config.DB.Raw(query, day, arrivalTime, departureTime, document).Row().Scan(&scheduleID)
// 		if err != nil {
// 			continue
// 		}
// 	}

// 	return c.JSON(http.StatusOK, "Carga masiva completada")
// }

func UploadExcelAndAssignSchedules(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Error al cargar el archivo")
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error al abrir el archivo subido")
	}
	defer src.Close()

	excelFile, err := excelize.OpenReader(src)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error al abrir el archivo Excel")
	}

	rows, err := excelFile.GetRows("horario")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error al leer las filas del archivo Excel")
	}

	for _, row := range rows {
		document := row[0]
		day := row[1]
		arrivalTime := row[2]
		departureTime := row[3]

		query := fmt.Sprintf(`INSERT INTO "schedules" ("day","arrival_time","departure_time","fk_collaborator_id")
                             SELECT ?, ?, ?, c."id"
                             FROM "collaborators" c
                             WHERE c."document" = ?
                             RETURNING "id"`)

		var scheduleID int
		err := config.DB.Raw(query, day, arrivalTime, departureTime, document).Row().Scan(&scheduleID)
		if err != nil {
			continue
		}
	}

	return c.JSON(http.StatusOK, "Carga masiva completada")
}
