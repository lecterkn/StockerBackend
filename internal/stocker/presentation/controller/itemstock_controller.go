package controller

import (
	"h11/backend/internal/stocker/application/service"
	"h11/backend/internal/stocker/common"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ItemStockController struct {
	itemStockService     service.ItemStockService
	authorizationService service.StoreAuthorizationService
}

// NewItemStockController /* プロバイダ
func NewItemStockController(itemStockService service.ItemStockService, authorizationService service.StoreAuthorizationService) ItemStockController {
	return ItemStockController{
		itemStockService,
		authorizationService,
	}
}

// Index /* 商品詳細一覧取得用エンドポイント
//
//	@Summary	商品詳細一覧取得
//	@Tags		item_stock
//	@Produce	json
//	@Param		storeId	path string		true	"店舗ID"
//	@Success	200		{object}	ItemStockListResponse{list=[]ItemStockResponse}
//	@Router		/stores/{storeId}/itemsStocks [get]
func (c ItemStockController) Index(ctx *fiber.Ctx) error {
	// ユーザーID取得
	userId, err := common.GetUserIdByContext(ctx)
	if err != nil {
		return err
	}
	// 店舗ID
	storeId, err := uuid.Parse(ctx.Params("storeId"))
	if err != nil {
		return err
	}
	// 店舗とユーザーの認証
	if err := c.authorizationService.IsUserRelated(storeId, *userId); err != nil {
		return err
	}
	// 商品詳細一覧取得
	listOutput, err := c.itemStockService.Index(service.ItemStockServiceQueryListInput{
		StoreId: storeId,
	})
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("internal error")
	}
	// レスポンスに変換
	var list []ItemStockResponse
	for _, output := range listOutput.List {
		list = append(list, ItemStockResponse(output))
	}
	return ctx.JSON(ItemStockListResponse{
		list,
	})
}

// Select /* 商品詳細取得用エンドポイント
//
//	@Summary	商品詳細取得
//	@Tags		item_stock
//	@Produce	json
//	@Param		storeId	path string		true	"店舗ID"
//	@Param		itemId	path string		true	"商品ID"
//	@Success	200		{object}	ItemStockResponse
//	@Router		/stores/{storeId}/items/{itemId}/stocks [get]
func (c ItemStockController) Select(ctx *fiber.Ctx) error {
	// ユーザーID取得
	userId, err := common.GetUserIdByContext(ctx)
	if err != nil {
		return err
	}
	// 店舗ID取得
	storeId, err := uuid.Parse(ctx.Params("storeId"))
	if err != nil {
		return err
	}
	// 商品ID取得
	id, err := uuid.Parse(ctx.Params("itemId"))
	if err != nil {
		return ctx.Status(http.StatusForbidden).SendString("invalid id")
	}
	// 店舗とユーザーの認証
	if err := c.authorizationService.IsUserRelated(storeId, *userId); err != nil {
		return err
	}
	// 商品詳細取得
	output, err := c.itemStockService.Select(service.ItemStockServiceQueryInput{
		StoreId: storeId,
		Id:      id,
	})
	if err != nil {
		return ctx.Status(http.StatusNotFound).SendString("not found")
	}
	return ctx.JSON(ItemStockResponse(*output))
}

// Create /* 商品詳細作成用エンドポイント
//
//	@Summary	商品詳細登録
//	@Tags		item_stock
//	@Produce	json
//	@Param		storeId	path string		true				"店舗ID"
//	@Param		itemId	path string		true				"商品ID"
//	@Param		request	body		ItemStockRequest	true	"商品詳細作成リクエスト"
//	@Success	200		{object}	ItemStockResponse
//	@Router		/stores/{storeId}/items/{itemId}/stocks [post]
func (c ItemStockController) Create(ctx *fiber.Ctx) error {
	// ユーザーID取得
	userId, err := common.GetUserIdByContext(ctx)
	if err != nil {
		return err
	}
	// 店舗ID取得
	storeId, err := uuid.Parse(ctx.Params("storeId"))
	if err != nil {
		return err
	}
	// 商品ID取得
	id, err := uuid.Parse(ctx.Params("itemId"))
	if err != nil {
		return ctx.Status(http.StatusForbidden).SendString("invalid id")
	}
	// 商品詳細作成リクエスト取得
	var request ItemStockRequest
	if err = ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusForbidden).SendString("invalid requestBody")
	}
	// 店舗とユーザーの認証
	if err := c.authorizationService.IsUserRelated(storeId, *userId); err != nil {
		return err
	}
	// 商品詳細作成
	output, err := c.itemStockService.Create(service.ItemStockServiceInput{
		StoreId:  storeId,
		ItemId:   id,
		Place:    request.Place,
		Price:    request.Price,
		Stock:    request.Stock,
		StockMin: request.StockMin,
	})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("internal error")
	}
	return ctx.JSON(ItemStockResponse(*output))
}

// Update /* 商品詳細更新
//
//	@Summary	商品詳細更新
//	@Tags		item_stock
//	@Produce	json
//	@Param		storeId	path string		true				"店舗ID"
//	@Param		itemId	path string		true				"商品ID"
//	@Param		request	body		ItemStockRequest	false	"商品詳細更新リクエスト"
//	@Success	200		{object}	ItemStockResponse
//	@Router		/stores/{storeId}/items/{itemId}/stocks [patch]
func (c ItemStockController) Update(ctx *fiber.Ctx) error {
	// ユーザーID取得
	userId, err := common.GetUserIdByContext(ctx)
	if err != nil {
		return err
	}
	// 店舗ID取得
	storeId, err := uuid.Parse(ctx.Params("storeId"))
	if err != nil {
		return err
	}
	// 商品ID取得
	id, err := uuid.Parse(ctx.Params("itemId"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("invalid id")
	}
	// 商品詳細更新リクエスト取得
	var request ItemStockRequest
	if err = ctx.BodyParser(&request); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("invalid requestBody")
	}
	// 店舗とユーザーの認証
	if err := c.authorizationService.IsUserRelated(storeId, *userId); err != nil {
		return err
	}
	// 商品詳細更新
	output, err := c.itemStockService.Update(service.ItemStockServiceInput{
		StoreId:  storeId,
		ItemId:   id,
		Place:    request.Place,
		Price:    request.Price,
		Stock:    request.Stock,
		StockMin: request.StockMin,
	})
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("internal error")
	}
	return ctx.JSON(ItemStockResponse(*output))
}

type ItemStockListResponse struct {
	List []ItemStockResponse `json:"list" validate:"required"`
}

type ItemStockResponse struct {
	ItemId    uuid.UUID `json:"itemId" validate:"required"`
	StoreId   uuid.UUID `json:"storeId" validate:"required"`
	Place     string    `json:"place" validate:"required"`
	Price     *int      `json:"price"`
	Stock     int       `json:"stock" validate:"required"`
	StockMin  int       `json:"stockMin" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

type ItemStockRequest struct {
	Place    string `json:"place" validate:"required"`
	Price    *int   `json:"price"`
	Stock    int    `json:"stock" validate:"required"`
	StockMin int    `json:"stockMin" validate:"required"`
}
