package controllers

import (
	"net/http"

	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	userEntity := entity.User{}
	err := c.Bind(&userEntity)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userModel, err := services.AuthenticateUser(userEntity.Email, userEntity.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := services.GenerateToken(userModel)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error generating token")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		"user":  userModel,
	})
}

func GetUserInfo(c echo.Context) error {
	user := c.Get("userToken").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"email":    claims["email"],
		"document": claims["document"],
		"fName":    claims["fName"],
		"lName":    claims["lName"],
		"rol":      claims["rol"],
		"roleName": claims["roleName"],
	})
}

func ValidateToken(c echo.Context) error {
	user := c.Get("userToken").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	rol, ok := claims["rol"]
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Error getting role")
	}
	return c.JSON(http.StatusOK, rol)
}

func Logout(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")

	if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	tokenString := authHeader[7:]
	services.InvalidateToken(tokenString)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Sesión cerrada con éxito",
	})
}
