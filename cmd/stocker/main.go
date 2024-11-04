package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"h11/backend/internal/stocker"
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

	e.GET("/items", controllerSets.ItemController.Index)
	e.POST("/items", controllerSets.ItemController.Create)

	e.Start(":")
}
