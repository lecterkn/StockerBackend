package controller

import (
	"time"

	"h11/backend/internal/stocker/application/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type StockOutController struct {
	stockOutUsecase usecase.StockOutUsecase
}

func NewStockOutController(stockOutUsecase usecase.StockOutUsecase) StockOutController {
	return StockOutController{
		stockOutUsecase,
	}
}

// Index /* 販売履歴一覧取得用エンドポイント
//
//	@Summary	販売履歴一覧取得取得
//	@Tags		stock_outs
//	@Produce	json
//	@Param		storeId	path		string	true	"店舗ID"
//	@Success	200		{object}	StockOutListResponse{list=[]StockOutResponse}
//	@Router		/stores/{storeId}/stockOuts [get]
func (c StockOutController) Index(ctx *fiber.Ctx) error {
	storeId, err := uuid.Parse(ctx.Params("storeId"))
	if err != nil {
		return err
	}
	outputList, err := c.stockOutUsecase.GetStockOuts(storeId)
	if err != nil {
		return err
	}
	var list []StockOutResponse
	for _, output := range outputList.List {
		list = append(list, StockOutResponse(output))
	}
	return ctx.JSON(StockOutListResponse{
		list,
	})
}

// Create /* 販売履歴作成用エンドポイント
//
//	@Summary	販売履歴作成取得
//	@Tags		stock_outs
//	@Produce	json
//	@Param		storeId	path		string					true	"店舗ID"
//	@Param		request	body		StockOutCreateRequest	true	"販売履歴作成リクエスト"
//	@Success	200		{object}	StockOutResponse
//	@Router		/stores/{storeId}/stockOuts [post]
func (c StockOutController) Create(ctx *fiber.Ctx) error {
	storeId, err := uuid.Parse(ctx.Params("storeId"))
	if err != nil {
		return err
	}
	var request StockOutCreateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}
	output, err := c.stockOutUsecase.CreateStockOut(storeId, usecase.StockOutUsecaseCommandInput(request))
	if err != nil {
		return err
	}
	return ctx.JSON(StockOutResponse(*output))
}

type StockOutResponse struct {
	Id        uuid.UUID `json:"id"        validate:"required"`
	StoreId   uuid.UUID `json:"storeId"   validate:"required"`
	ItemId    uuid.UUID `json:"itemId"    validate:"required"`
	Place     *string   `json:"place"     validate:"required"`
	Price     int       `json:"price"     validate:"required"`
	Stocks    int       `json:"stocks"    validate:"required"`
	Name      string    `json:"name"      validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

type StockOutCreateRequest struct {
	ItemId uuid.UUID `json:"itemId" validate:"required"`
	Place  *string   `json:"place"`
	Price  int       `json:"price" validate:"required"`
	Stocks int       `json:"stocks" validate:"required"`
}

type StockOutListResponse struct {
	List []StockOutResponse `validate:"required"`
}
