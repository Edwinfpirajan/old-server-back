package router

import (
	"github.com/DevEdwinF/smartback.git/internal/app/controllers"
	"github.com/DevEdwinF/smartback.git/internal/config/middleware"
	"github.com/labstack/echo/v4"
)

func CollaboratorRoutes(e *echo.Echo) {

	group := e.Group("/api/collaborator")

	// group.POST("/save", controller.SaveCollaborator)
	group.GET("/all", controllers.GetAllCollaboratorsController, middleware.AuthToken)
	group.GET("/find/:document", controllers.GetCollaborator, middleware.AuthToken)
	group.DELETE("/delete/:doc", controllers.DeleteCollaborator, middleware.AuthToken)
	group.GET("/test", controllers.GetAllColab)
}
