package roles

import (
	"github.com/DevEdwinF/smartback.git/internal/app/controllers"
	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/labstack/echo/v4"
)

func RolesRoutes(e *echo.Echo) {
	rolesService := services.NewRolesService()
	rolesController := controllers.NewRolesController(rolesService)

	group := e.Group("/api/roles")
	group.GET("/all", rolesController.GetAllRoles)
}
