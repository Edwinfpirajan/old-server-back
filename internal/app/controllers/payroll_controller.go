package controllers

import (
	"net/http"

	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/labstack/echo/v4"
)

type PayrollController struct {
	Service *services.PayrollService
}

func NewPayrollController(s *services.PayrollService) *PayrollController {
	return &PayrollController{
		Service: s,
	}
}

func (controller *PayrollController) GetAllPayRoll(c echo.Context) error {
	payroll, err := controller.Service.GetAllPayRollService()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, payroll)
}
