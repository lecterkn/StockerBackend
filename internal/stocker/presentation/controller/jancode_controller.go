package controller

import (
	"h11/backend/internal/stocker/application/service"

	"github.com/gofiber/fiber/v2"
)

type JancodeController struct {
    jancodeService service.JancodeService
}

func NewJancodeController(jancodeService service.JancodeService) JancodeController {
    return JancodeController{
        jancodeService,
    }
}

// Select /* 製品情報取得エンドポイント
//	@Summary	製品情報取得
//	@Tags		jancode
//	@Produce	json
//	@Param		janCode	path string		true	"JANコード"
//	@Success	200		{object}	JancodeResponse
//	@Router		/products/{janCode} [get]
func (c JancodeController) Select(ctx *fiber.Ctx) error {
    janCode := ctx.Params("janCode")
    output, err := c.jancodeService.GetProductByCode(janCode)
    if err != nil {
        return err
    }
    return ctx.JSON(JancodeResponse(*output))
}

type JancodeResponse struct {
    Name string `json:"name"`
    BrandName string `json:"brandName"`
    MakerName string `json:"makerName"`
}
