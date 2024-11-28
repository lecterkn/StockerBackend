package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"h11/backend/internal/stocker/application/service"
	"net/http"
	"time"
)

type ItemController struct {
	itemService service.ItemService
}

// NewItemController /* アイテムコントローラーのプロバイダ
func NewItemController(itemService service.ItemService) ItemController {
	return ItemController{
		itemService,
	}
}

// Index /* アイテムを一覧取得
//	@Summary	商品一覧取得
//	@Tags		item
//	@Produce	json
//	@Success	200	{object}	ItemListResponse{list=[]ItemResponse}
//	@Router		/items [get]
func (c ItemController) Index(ctx *fiber.Ctx) error {
	listOutput, err := c.itemService.GetItems()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString("internal error")
	}
	var list []ItemResponse
	for _, output := range listOutput.List {
		list = append(list, ItemResponse(output))
	}
	return ctx.JSON(ItemListResponse{
		List: list,
	})
}

// Create /* アイテムを新規作成
//	@Summary	商品新規作成
//	@Tags		item
//	@Produce	json
//	@Param		request	body		ItemRequest false "アイテム作成リクエスト"
//	@Success	200		{object}	ItemResponse
//	@Router		/items [post]
func (c ItemController) Create(ctx *fiber.Ctx) error {
	var request ItemRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}
	output, err := c.itemService.CreateItem(service.ItemServiceInput(request))
	if err != nil {
		return err
	}
	return ctx.JSON(ItemResponse(*output))
}

// Update /* アイテムを更新
//	@Summary	商品更新
//	@Tags		item
//	@Produce	json
//	@Param		request	body		ItemRequest false "アイテム作成リクエスト"
//	@Success	200		{object}	ItemResponse
//	@Router		/items [patch]
func (c ItemController) Update(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return err
	}
	var request ItemRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}
	output, err := c.itemService.UpdateItem(service.ItemServiceUpdateInput{
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
	Name      string    `json:"name"`
	JanCode   string    `json:"janCode"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ItemListResponse struct {
	List []ItemResponse `json:"list"`
}
