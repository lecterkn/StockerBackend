package controller

import (
	"net/http"
	"time"

	"h11/backend/internal/stocker/application/usecase"
	"h11/backend/internal/stocker/common"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ItemController struct {
	itemUsecase          usecase.ItemUsecase
	authorizationUsecase usecase.StoreAuthorizationUsecase
}

// NewItemController /* アイテムコントローラーのプロバイダ
func NewItemController(itemUsecase usecase.ItemUsecase, authorizationUsecase usecase.StoreAuthorizationUsecase) ItemController {
	return ItemController{
		itemUsecase,
		authorizationUsecase,
	}
}

// Index /* 商品の一覧取得
//
//	@Summary	商品一覧取得
//	@Tags		item
//	@Produce	json
//	@Param		storeId	path		string	true	"店舗ID"
//	@Param		name	query		string	false	"商品名"
//	@Param		janCode	query		string	false	"Janコード"
//	@Success	200		{object}	ItemListResponse{list=[]ItemResponse}
//	@Router		/stores/{storeId}/items [get]
func (c ItemController) Index(ctx *fiber.Ctx) error {
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
	// クエリパラメータを取得
	var query ItemListQuery
	if err := ctx.QueryParser(&query); err != nil {
		return err
	}
	// 店舗とユーザーの認証
	if err := c.authorizationUsecase.IsUserRelated(storeId, *userId); err != nil {
		return err
	}
	// アイテムを取得
	listOutput, err := c.itemUsecase.GetItems(storeId, query.jancode, query.name)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("internal error")
	}
	// レスポンス変換
	var list []ItemResponse
	for _, output := range listOutput.List {
		list = append(list, ItemResponse(output))
	}
	return ctx.JSON(ItemListResponse{
		List: list,
	})
}

// Create /* 商品を新規作成
//
//	@Summary	商品新規作成
//	@Tags		item
//	@Produce	json
//	@Param		request	body		ItemRequest	true	"商品作成リクエスト"
//	@Param		storeId	path		string		true	"店舗ID"
//	@Success	200		{object}	ItemResponse
//	@Router		/stores/{storeId}/items [post]
func (c ItemController) Create(ctx *fiber.Ctx) error {
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
	// 商品作成リクエストを取得
	var request ItemRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}
	// 店舗とユーザーの認証
	if err := c.authorizationUsecase.IsUserRelated(storeId, *userId); err != nil {
		return err
	}
	// 商品作成
	output, err := c.itemUsecase.CreateItem(usecase.ItemUsecaseInput{
		StoreId: storeId,
		Name:    request.Name,
		JanCode: request.JanCode,
	})
	if err != nil {
		return err
	}
	return ctx.JSON(ItemResponse(*output))
}

// Update /* 商品を更新
//
//	@Summary	商品更新
//	@Tags		item
//	@Produce	json
//	@Param		storeId	path		string		true	"店舗ID"
//	@Param		itemId	path		string		true	"商品ID"
//	@Param		request	body		ItemRequest	true	"商品更新リクエスト"
//	@Success	200		{object}	ItemResponse
//	@Router		/stores/{storeId}/items/{itemId} [patch]
func (c ItemController) Update(ctx *fiber.Ctx) error {
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
		return err
	}
	// 商品更新リクエスト取得
	var request ItemRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}
	// 店舗とユーザーの認証
	if err := c.authorizationUsecase.IsUserRelated(storeId, *userId); err != nil {
		return err
	}
	// 商品更新
	output, err := c.itemUsecase.UpdateItem(usecase.ItemUsecaseUpdateInput{
		StoreId: storeId,
		Id:      id,
		Name:    request.Name,
		JanCode: request.JanCode,
	})
	if err != nil {
		return err
	}
	return ctx.JSON(ItemResponse(*output))
}

type ItemListQuery struct {
	name    *string `query:"name"`
	jancode *string `query:"janCode"`
}

type ItemRequest struct {
	Name    string `json:"name" validate:"required"`
	JanCode string `json:"janCode" validate:"required"`
}

type ItemResponse struct {
	Id        uuid.UUID `json:"id" validate:"required"`
	StoreId   uuid.UUID `json:"storeId" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	JanCode   string    `json:"janCode" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

type ItemListResponse struct {
	List []ItemResponse `json:"list" validate:"required"`
}
