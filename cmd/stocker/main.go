package main

import (
	"fmt"
	"h11/backend/internal/stocker"
	"h11/backend/internal/stocker/application/service"
	"h11/backend/internal/stocker/infrastructure/database"
	"h11/backend/internal/stocker/infrastructure/implements"
	"h11/backend/internal/stocker/presentation/controller"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// .env 読み込み
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// echo作成
	e := echo.New()

	// DI
	controllerSets := stocker.InitializeController()

	e.GET("/items", controllerSets.ItemController.GetItems)

	e.Start(":")
}
