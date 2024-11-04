package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"h11/backend/internal/stocker"
	"os"
)

func main() {
	// .env 読み込み
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	port, ok := os.LookupEnv("ECHO_SERVER_PORT")
	if !ok {
		fmt.Println("\"ECHO_SERVER_PORT\" is not set")
		return
	}

	// echo作成
	e := echo.New()

	// DI
	controllerSets := stocker.InitializeController()

	e.GET("/items", controllerSets.ItemController.Index)
	e.POST("/items", controllerSets.ItemController.Create)

	e.Start(":" + port)
}
