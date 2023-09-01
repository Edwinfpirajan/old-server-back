package controllers

import (
	"net/http"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (h *UserController) CreateUser(c echo.Context) error {
	var user entity.UserData
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	err := h.userService.CreateUser(user)
	if err != nil {
		switch err.Error() {
		case "El usuario ya existe":
			return echo.NewHTTPError(http.StatusBadRequest, "El usuario ya existe")
		case "failed to create user":
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to create user")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
		}
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserController) GetAllUsers(c echo.Context) error {
	var users []models.User
	err := h.userService.GetAllUsers(&users)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get all users")
	}

	return c.JSON(http.StatusOK, users)
}

func (h *UserController) GetUserById(c echo.Context) error {
	document := c.Param("doc")
	user, err := h.userService.GetUserById(document)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user")
	}

	return c.JSON(http.StatusOK, user)
}

// func (h *UserController) UpdateUser(c echo.Context) error {
// 	var user entity.UserData
// 	if err := c.Bind(&user); err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
// 	}

// 	err := h.userService.UpdateUser(user)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update user")
// 	}
// 	return c.JSON(http.StatusOK, map[string]string{
// 		"message": "Usuario actualizado exitosamente",
// 	})

// }

func (h *UserController) UpdateUser(c echo.Context) error {
	var user entity.UserData
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}

	// newFkRoleID := user.FkRoleId

	err := h.userService.UpdateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to update user")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Usuario actualizado exitosamente",
	})
}

func (h *UserController) DeleteUser(c echo.Context) error {
	document := c.Param("doc")
	err := h.userService.DeleteUser(document)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to delete user")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Usuario eliminado",
	})
}
