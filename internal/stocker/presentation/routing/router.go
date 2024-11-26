package routing

import (
	"h11/backend/internal/stocker"

	"github.com/gofiber/fiber/v3"
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
}
