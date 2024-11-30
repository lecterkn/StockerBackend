package controller

import (
	"h11/backend/internal/stocker/application/service"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
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

// Index /* 商品詳細一覧取得用エンドポイント
//	@Summary	商品詳細一覧取得
//	@Tags		item_stock
//	@Produce	json
//	@Success	200	{object}	ItemStockListResponse{list=[]ItemStockResponse}
//	@Router		/itemsStocks [get]
func (c ItemStockController) Index(ctx *fiber.Ctx) error {
	listOutput, err := c.itemStockService.Index()
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("internal error")
	}
	var list []ItemStockResponse
	for _, output := range listOutput.List {
		list = append(list, ItemStockResponse(output))
	}
	return ctx.JSON(ItemStockListResponse{
		list,
	})
}

// Select /* 商品詳細取得用エンドポイント
//	@Summary	商品詳細取得
//	@Tags		item_stock
//	@Produce	json
//	@Success	200	{object}	ItemStockResponse
//	@Router		/items/{itemId}/stocks [get]
func (c ItemStockController) Select(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("itemId"))
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
//	@Summary	商品詳細登録
//	@Tags		item_stock
//	@Produce	json
//	@Param		request	body		ItemStockRequest	false	"商品詳細作成リクエスト"
//	@Success	200		{object}	ItemStockResponse
//	@Router		/items/{itemId}/stocks [post]
func (c ItemStockController) Create(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("itemId"))
	if err != nil {
		return ctx.Status(http.StatusForbidden).SendString("invalid id")
	}
	var request ItemStockRequest
	if err = ctx.BodyParser(&request); err != nil {
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

// Update /* 商品詳細更新
//	@Summary	商品詳細更新
//	@Tags		item_stock
//	@Produce	json
//	@Param		request	body		ItemStockRequest	false	"商品詳細更新リクエスト"
//	@Success	200		{object}	ItemStockResponse
//	@Router		/items/{itemId}/stocks [patch]
func (c ItemStockController) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("itemId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("invalid id")
	}
	var request ItemStockRequest
	if err = ctx.BodyParser(&request); err != nil {
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

type ItemStockListResponse struct {
	List []ItemStockResponse `json:"list"`
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
