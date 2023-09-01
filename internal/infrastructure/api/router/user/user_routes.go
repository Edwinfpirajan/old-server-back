package user

import (
	"github.com/DevEdwinF/smartback.git/internal/app/controllers"
	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {

	userService := &services.UserService{}
	UserController := controllers.NewUserController(userService)

	group := e.Group("/api/user")
	group.POST("/create", UserController.CreateUser)
	group.GET("/all", UserController.GetAllUsers)
	group.GET("/:doc", UserController.GetUserById)
	group.PATCH("/update", UserController.UpdateUser)
	group.DELETE("/delete/:doc", UserController.DeleteUser)
}
