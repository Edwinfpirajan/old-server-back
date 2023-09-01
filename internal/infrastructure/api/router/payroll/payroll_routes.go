package payroll

import (
	"github.com/DevEdwinF/smartback.git/internal/app/controllers"
	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/labstack/echo/v4"
)

func PayrollRouter(e *echo.Echo) {
	payrollService := services.NewPayrollService()
	payrollController := controllers.NewPayrollController(payrollService)

	group := e.Group("/api/payroll")
	group.GET("/hearquarters/all", payrollController.GetAllPayRoll /* middleware.AuthToken */)
}
