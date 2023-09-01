package controllers

import (
	"net/http"

	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/labstack/echo/v4"
)

// func GetTest(c echo.Context) error {
// 	colaborador := []BiEmple{}

// 	config.KDB.Table("bi_emple").Select("*").Scan(&colaborador)
// 	return c.JSON(http.StatusOK, colaborador)
// }

func GetAllColab(c echo.Context) error {
	collaborators, err := services.GetAllColab()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, collaborators)
}

func GetColab(c echo.Context) error {
	document := c.Param("document")

	collaborator, err := services.GetColabById(document)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, collaborator)
}
