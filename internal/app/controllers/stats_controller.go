package controllers

import (
	"net/http"
	"time"

	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/labstack/echo/v4"
)

func CountAttendancesAll(c echo.Context) error {
	count, err := services.CountAttendances()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]int64{"count": count})
}

func CountAttendanceDay(c echo.Context) error {
	today := time.Now()
	count, err := services.CountAttendanceForDay(today)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]int64{"count": count})
}

func CountLateAttendancesForDay(c echo.Context) error {
	today := time.Now()
	count, err := services.CountAttendanceForDayByLate(today, true)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]int64{"count": count})
}

func CountOnTimeAttendancesForDay(c echo.Context) error {
	today := time.Now()
	count, err := services.CountAttendanceForDayByLate(today, false)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]int64{"count": count})
}

func TotalCollaboratorsActiveController(c echo.Context) error {
	count, err := services.TotalCollaboratorsActiveService()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]int64{"count": count})
}
