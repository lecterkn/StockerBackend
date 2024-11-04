package controller

import (
	"github.com/google/uuid"
	"h11/backend/internal/stocker/application/service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type ItemController struct {
	ItemService service.ItemService
}

func NewItemController(itemService service.ItemService) ItemController {
	return ItemController{
		ItemService: itemService,
	}
}

func (c ItemController) Index(ctx echo.Context) error {
	listOutput, err := c.ItemService.GetItems()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "internal error")
	}
	var list []ItemResponse
	for _, output := range listOutput.List {
		list = append(list, ItemResponse(output))
	}
	return ctx.JSON(http.StatusOK, ItemListResponse{
		List: list,
	})
}

func (c ItemController) Create(ctx echo.Context) error {
	var request ItemCreateRequest
	if err := ctx.Bind(&request); err != nil {
		return err
	}
	output := c.ItemService.CreateItem(service.ItemServiceInput(request))
	return ctx.JSON(http.StatusOK, ItemResponse(*output))
}

type ItemCreateRequest struct {
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
