package routing

import (
	"h11/backend/internal/stocker"
	"h11/backend/internal/stocker/common"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// swagger handler
	"github.com/gofiber/swagger"
	// jwtware
	_ "h11/backend/docs"

	jwtware "github.com/gofiber/contrib/jwt"
)

func SetRouting(f *fiber.App) {
	// di
	controllerSets := stocker.InitializeController()

	// CORS
	setCors(f)

	// Swaggger
	setSwagger(f)

	// UserController
	f.Post("/register", controllerSets.UserController.Create)
	f.Post("/login", controllerSets.UserController.Login)

	// JWT Middleware
	// このコードより後述のエンドポイントにJWT認証が適応される
	f.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    common.GetJwtSecretKey(),
		},
	}))

	// JancodeController
	f.Get("/products/:janCode", controllerSets.JancodeController.Select)

	// ItemController
	f.Get("/stores/:storeId/items", controllerSets.ItemController.Index)
	f.Post("/stores/:storeId/items", controllerSets.ItemController.Create)
	f.Patch("/stores/:storeId/items/:itemId", controllerSets.ItemController.Update)

	// ItemStockController
	f.Get("/stores/:storeId/itemStocks", controllerSets.ItemStockController.Index)
	f.Get("/stores/:storeId/items/:itemId/stocks", controllerSets.ItemStockController.Select)
	f.Post("/stores/:storeId/items/:itemId/stocks", controllerSets.ItemStockController.Create)
	f.Patch("/stores/:storeId/items/:itemId/stocks", controllerSets.ItemController.Update)

	// StoreController
	f.Get("/stores", controllerSets.StoreController.Index)
	f.Post("/stores", controllerSets.StoreController.Create)
	f.Patch("/stores/:storeId", controllerSets.StoreController.Update)

	// StockInController
	f.Get("/stores/:storeId/stockIns", controllerSets.StockInController.Index)
	f.Post("/stores/:storeId/stockIns", controllerSets.StockInController.Create)
}

func setCors(f *fiber.App) {
	f.Use(cors.New(
		cors.Config{
			AllowOrigins: "http://localhost:5173",
			AllowHeaders: "Authorization, Content-Type, Origin, Accept",
			AllowMethods: "GET, POST, PATCH, PUT, DELETE, OPTIONS",
		}))
}

func setSwagger(f *fiber.App) {
	f.Get("/swagger/*", swagger.HandlerDefault) // default
}
