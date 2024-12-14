package controller

import (
	"h11/backend/internal/stocker/application/usecase"

	"github.com/gofiber/fiber/v2"
)

type JancodeController struct {
	jancodeUsecase usecase.JancodeUsecase
}

func NewJancodeController(jancodeUsecase usecase.JancodeUsecase) JancodeController {
	return JancodeController{
		jancodeUsecase,
	}
}

// Select /* 製品情報取得エンドポイント
//
//	@Summary	製品情報取得
//	@Tags		jancode
//	@Produce	json
//	@Param		janCode	path string		true	"JANコード"
//	@Success	200		{object}	JancodeResponse
//	@Router		/products/{janCode} [get]
func (c JancodeController) Select(ctx *fiber.Ctx) error {
	janCode := ctx.Params("janCode")
	output, err := c.jancodeUsecase.GetProductByCode(janCode)
	if err != nil {
		return err
	}
	return ctx.JSON(JancodeResponse(*output))
}

type JancodeResponse struct {
	Name      string `json:"name" validate:"required"`
	BrandName string `json:"brandName" validate:"required"`
	MakerName string `json:"makerName" validate:"required"`
}
