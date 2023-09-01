package main

import (
	"fmt"
	"os"
	"strconv"

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

	// e.Logger.Fatal(e.Start(":8080"))
	PORT, _ := strconv.Atoi(os.Getenv("PORT"))
	HOST := os.Getenv("SERVER_HOST")

	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%d", HOST, PORT)))
	//testint
	//testing2
}
