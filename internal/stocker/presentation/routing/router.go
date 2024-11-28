package routing

import (
	"h11/backend/internal/stocker"

	"github.com/gofiber/fiber/v2"
    "github.com/gofiber/swagger" // swagger handler
    _ "h11/backend/docs"
)

func SetRouting(f *fiber.App) {
	// di
	controllerSets := stocker.InitializeController()

	// ItemController
	f.Get("/items", controllerSets.ItemController.Index)
	f.Post("/items", controllerSets.ItemController.Create)
	f.Patch("/items/:id", controllerSets.ItemController.Update)

	// ItemStockController
	f.Get("/items/:id/stocks", controllerSets.ItemStockController.Select)
	f.Post("/items/:id/stocks", controllerSets.ItemStockController.Create)

    setSwagger(f)
}

func setSwagger(f *fiber.App) {
    f.Get("/swagger/*", swagger.HandlerDefault) // default
}
