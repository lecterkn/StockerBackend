package main

import (
	"fmt"
	"h11/backend/internal/stocker/presentation/routing"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
    _ "h11/backend/docs"
)

//	@title			StockerAPI
//	@version		1.0
//	@description	stocker application server.
//	@host			localhost:8810
//	@BasePath		/
func main() {
	// .env 読み込み
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
		return
	}

	// port番号 読み込み
	port, ok := os.LookupEnv("FIBER_SERVER_PORT")
	if !ok {
		fmt.Println("\"FIBER_SERVER_PORT\" is not set")
		return
	}

	// fiber作成
	f := fiber.New()

	// routing
	routing.SetRouting(f)

	// start echo
	f.Listen(":" + port)
}
