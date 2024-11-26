package controller

import (
	"h11/backend/internal/stocker/application/service"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type ItemStockController struct {
	itemStockService service.ItemStockService
}

// NewItemStockController /* プロバイダ
func NewItemStockController(itemStockService service.ItemStockService) ItemStockController {
	return ItemStockController{
		itemStockService,
	}
}

// Select /* 商品詳細取得用エンドポイント
func (c ItemStockController) Select(ctx fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusForbidden).SendString("invalid id")
	}
	output, err := c.itemStockService.Select(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).SendString("not found")
	}
	return ctx.JSON(ItemStockResponse(*output))
}

// Create /* 商品詳細作成用エンドポイント
func (c ItemStockController) Create(ctx fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusForbidden).SendString("invalid id")
	}
	var request ItemStockRequest
	if err = ctx.Bind().Body(&request); err != nil {
		return ctx.Status(http.StatusForbidden).SendString("invalid requestBody")
	}
	output, err := c.itemStockService.Create(service.ItemStockServiceInput{
		ItemId:   id,
		Place:    request.Place,
		Stock:    request.Stock,
		StockMin: request.StockMin,
	})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("internal error")
	}
	return ctx.JSON(ItemStockResponse(*output))
}

func (c ItemStockController) Update(ctx fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("invalid id")
	}
	var request ItemStockRequest
	if err = ctx.Bind().Body(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("invalid requestBody")
	}
	output, err := c.itemStockService.Update(service.ItemStockServiceInput{
		ItemId:   id,
		Place:    request.Place,
		Stock:    request.Stock,
		StockMin: request.StockMin,
	})
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("internal error")
	}
	return ctx.JSON(ItemStockResponse(*output))
}

type ItemStockResponse struct {
	ItemId    uuid.UUID `json:"item_id"`
	Place     string    `json:"place"`
	Stock     int       `json:"stock"`
	StockMin  int       `json:"stock_min"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ItemStockRequest struct {
	Place    string `json:"place"`
	Stock    int    `json:"stock"`
	StockMin int    `json:"stock_min"`
}
