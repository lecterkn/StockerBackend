package routing

import (
	"h11/backend/internal/stocker"
	"h11/backend/internal/stocker/common"

	"github.com/gofiber/fiber/v2"
	// swagger handler
	"github.com/gofiber/swagger"
	// jwtware
	_ "h11/backend/docs"

	jwtware "github.com/gofiber/contrib/jwt"
)

func SetRouting(f *fiber.App) {
	// di
	controllerSets := stocker.InitializeController()

	// Swaggger
    setSwagger(f)

	// UserController
	f.Post("/register", controllerSets.UserController.Create)
	f.Post("/login", controllerSets.UserController.Login)

	// JWT Middleware
	// このコードより後述のエンドポイントに適応される
	f.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key: common.GetJwtSecretKey(),
		},
	}))

	// ItemController
	f.Get("/items", controllerSets.ItemController.Index)
	f.Post("/items", controllerSets.ItemController.Create)
	f.Patch("/items/:itemId", controllerSets.ItemController.Update)

	// ItemStockController
	f.Get("/itemStocks", controllerSets.ItemStockController.Index)
	f.Get("/items/:itemId/stocks", controllerSets.ItemStockController.Select)
	f.Post("/items/:itemId/stocks", controllerSets.ItemStockController.Create)
	f.Patch("/items/:itemId/stocks", controllerSets.ItemController.Update)

	// StoreController
	f.Get("/stores", controllerSets.StoreController.Index)
	f.Post("/stores", controllerSets.StoreController.Create)
	f.Patch("/stores/:storeId", controllerSets.StoreController.Update)
}

func setSwagger(f *fiber.App) {
    f.Get("/swagger/*", swagger.HandlerDefault) // default
}
