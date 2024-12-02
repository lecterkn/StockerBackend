package controller

import (
	"h11/backend/internal/stocker/application/service"
	"h11/backend/internal/stocker/common"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ItemController struct {
	itemService service.ItemService
	authorizationService service.StoreAuthorizationService
}

// NewItemController /* アイテムコントローラーのプロバイダ
func NewItemController(itemService service.ItemService, authorizationService service.StoreAuthorizationService) ItemController {
	return ItemController{
		itemService,
		authorizationService,
	}
}

// Index /* 商品の一覧取得
//	@Summary	商品一覧取得
//	@Tags		item
//	@Produce	json
//	@Param		storeId	path		string	true	"店舗ID"
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
	// 店舗とユーザーの認証
	if err := c.authorizationService.IsUserRelated(storeId, *userId); err != nil {
		return err
	}
	// アイテムを取得
	listOutput, err := c.itemService.GetItems(service.ItemServiceQueryListInput{
		StoreId: storeId,
	})
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
	if err := c.authorizationService.IsUserRelated(storeId, *userId); err != nil {
		return err
	}
	// 商品作成
	output, err := c.itemService.CreateItem(service.ItemServiceInput{
		StoreId: storeId,
		Name: request.Name,
		JanCode: request.JanCode,
	})
	if err != nil {
		return err
	}
	return ctx.JSON(ItemResponse(*output))
}

// Update /* 商品を更新
//	@Summary	商品更新
//	@Tags		item
//	@Produce	json
//	@Param		storeId	path string		true		"店舗ID"
//	@Param		itemId	path string		true		"商品ID"
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
	if err := c.authorizationService.IsUserRelated(storeId, *userId); err != nil {
		return err
	}
	// 商品更新
	output, err := c.itemService.UpdateItem(service.ItemServiceUpdateInput{
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

type ItemRequest struct {
	Name    string `json:"name"`
	JanCode string `json:"janCode"`
}

type ItemResponse struct {
	Id        uuid.UUID `json:"id"`
	StoreId 	uuid.UUID `json:"storeId"`
	Name      string    `json:"name"`
	JanCode   string    `json:"janCode"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ItemListResponse struct {
	List []ItemResponse `json:"list"`
}
