package controller

import (
	"h11/backend/internal/stocker/application/service"
	"net/http"

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

func (c ItemController) GetItems(ctx echo.Context) error {
	items, err := c.ItemService.GetItems()
	if err != nil {
		return ctx.String(http.StatusInternalServerError, "internal error")
	}
	return ctx.JSON(http.StatusOK, items)
}
