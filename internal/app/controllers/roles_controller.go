package controllers

import (
	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/labstack/echo/v4"
)

type RolesController struct {
	Service *services.RolesService
}

func NewRolesController(s *services.RolesService) *RolesController {
	return &RolesController{
		Service: s,
	}
}

func (rc *RolesController) GetAllRoles(c echo.Context) error {
	roles, err := rc.Service.GetAllRoles()
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	return c.JSON(200, roles)
}
