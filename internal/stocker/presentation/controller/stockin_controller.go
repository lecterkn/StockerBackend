package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"h11/backend/internal/stocker/application/usecase"
)

type StockInController struct {
	stockInUsecase usecase.StockInUsecase
}

func NewStockInController(stockInUsecase usecase.StockInUsecase) StockInController {
	return StockInController{
		stockInUsecase,
	}
}

// Index /* 入荷履歴一覧取得用エンドポイント
//
//	@Summary	入荷履歴一覧取得
//	@Tags		stock_ins
//	@Produce	json
//	@Param		storeId	path		string	true	"店舗ID"
//	@Success	200		{object}	StockInListResponse
//	@Router		/stores/{storeId}/stockIns [get]
func (c StockInController) Index(ctx *fiber.Ctx) error {
	storeId, err := uuid.Parse(ctx.Params("storeId"))
	if err != nil {
		return err
	}
	listOutput, err := c.stockInUsecase.GetStockIns(storeId)
	if err != nil {
		return err
	}
	var list []StockInResponse
	for _, output := range listOutput.List {
		list = append(list, StockInResponse(output))
	}
	return ctx.JSON(StockInListResponse{
		list,
	})
}

// Create /* 入荷履歴作成用エンドポイント
//
//	@Summary	入荷履歴作成取得
//	@Tags		stock_ins
//	@Produce	json
//	@Param		storeId	path		string					true	"店舗ID"
//	@Param		request	body		StockInCreateRequest	true	"入荷履歴作成リクエスト"
//	@Success	200		{object}	StockInResponse
//	@Router		/stores/{storeId}/stockIns [post]
func (c StockInController) Create(ctx *fiber.Ctx) error {
	storeId, err := uuid.Parse(ctx.Params("storeId"))
	if err != nil {
		return err
	}
	var request StockInCreateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}
	output, err := c.stockInUsecase.CreateStockIn(storeId, usecase.StockInCommandInput(request))
	if err != nil {
		return err
	}
	return ctx.JSON(*output)
}

type StockInCreateRequest struct {
	ItemId uuid.UUID `json:"itemId" validate:"required"`
	Price  int       `json:"price"  validate:"required"`
	Stocks int       `json:"stocks" validate:"required"`
}

type StockInResponse struct {
	Id        uuid.UUID `json:"id"        validate:"required"`
	StoreId   uuid.UUID `json:"storeId"   validate:"required"`
	ItemId    uuid.UUID `json:"itemId"    validate:"required"`
	Name      string    `json:"name"      validate:"required"`
	Price     int       `json:"price"     validate:"required"`
	Stocks    int       `json:"stocks"    validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

type StockInListResponse struct {
	List []StockInResponse `json:"list" validate:"required"`
}
