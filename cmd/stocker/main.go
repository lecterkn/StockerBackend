package main

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
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

	// port番号 読み込み
	port, ok := os.LookupEnv("ECHO_SERVER_PORT")
	if !ok {
		fmt.Println("\"ECHO_SERVER_PORT\" is not set")
		return
	}

	// fiber作成
	f := fiber.New()

	// routing
	setRouting(f)
	// start echo
	f.Listen(":" + port)
}

// setRouting /* ルーティング設定
func setRouting(f *fiber.App) {
	// di
	controllerSets := stocker.InitializeController()

	f.Get("/items", controllerSets.ItemController.Index)
	f.Post("/items", controllerSets.ItemController.Create)
	f.Patch("/items/:id", controllerSets.ItemController.Update)
}
