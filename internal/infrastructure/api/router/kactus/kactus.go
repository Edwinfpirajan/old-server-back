package kactus

import (
	"github.com/DevEdwinF/smartback.git/internal/app/controllers"
	"github.com/labstack/echo/v4"
)

func KactusRouter(e *echo.Echo) {

	group := e.Group("/api/kactus")

	group.GET("/collaborators/all", controllers.GetAllColab)
	group.GET("/collaborators/:document", controllers.GetColab)
}
