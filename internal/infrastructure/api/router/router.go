package router

import (
	attendance "github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/attendance"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/auth"
	collaborator "github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/collaborator"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/kactus"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/payroll"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/roles"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/schedule"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/stats"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GlobalRouter(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		//toreto
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	auth.AuthRoutes(e)
	attendance.AttendanceRoutes(e)
	collaborator.CollaboratorRoutes(e)
	schedule.ScheduleRouter(e)
	stats.StatsRoutes(e)
	kactus.KactusRouter(e)
	user.UserRoutes(e)
	roles.RolesRoutes(e)
	payroll.PayrollRouter(e)

}
