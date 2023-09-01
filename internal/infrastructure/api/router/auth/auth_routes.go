package auth

import (
	"github.com/DevEdwinF/smartback.git/internal/app/controllers"
	"github.com/DevEdwinF/smartback.git/internal/config/middleware"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	group := e.Group("/auth")
	group.POST("/login", controllers.Login)
	group.POST("/logout", controllers.Logout, middleware.AuthToken)
	group.GET("/user-info", controllers.GetUserInfo, middleware.AuthToken)
	group.GET("/validate-token", controllers.ValidateToken, middleware.AuthToken)
}
