package main

import (
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDB()
	// config.KactusDB()
	e := echo.New()
	// services.RunCronJob()

	router.GlobalRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
	//testint
	//testing2
}
