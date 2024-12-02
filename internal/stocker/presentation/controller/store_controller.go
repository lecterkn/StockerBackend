package controller

import (
	"h11/backend/internal/stocker/application/service"
	"h11/backend/internal/stocker/common"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type StoreController struct {
	storeService service.StoreService
}

func NewStoreController(storeService service.StoreService) StoreController {
    return StoreController{
        storeService,
    }
}

// Index /* 店舗を一覧取得
//	@Summary	店舗一覧取得
//	@Tags		store
//	@Produce	json
//	@Success	200	{object}	StoreListResponse{list=[]StoreResponse}
//	@Router		/stores [get]
func (c StoreController) Index(ctx *fiber.Ctx) error {
    // JWTトークンからユーザーIDを取得
    userId, err := common.GetUserIdByContext(ctx)
    if err != nil {
        return err
    }
    // 店舗一覧取得
    listOutput, err :=c.storeService.Index(service.StoreServiceQueryListInput{
        UserId: *userId,
    })
    if err != nil {
        return err
    }
    // レスポンスに変換
    var list []StoreResponse
    for _, output := range(listOutput.List) {
        list = append(list, *c.toResponse(&output))
    }
    return ctx.JSON(StoreListResponse{
        List: list,
    })
}

// Create /* 店舗を作成
//	@Summary	店舗新規作成
//	@Tags		store
//	@Produce	json
//	@Param		request	body		StoreCreateRequest	false	"店舗新規作成リクエスト"
//	@Success	200		{object}	StoreResponse
//	@Router		/stores [post]
func (c StoreController) Create(ctx *fiber.Ctx) error {
    // JWTトークンからユーザーIDを取得
    userId, err := common.GetUserIdByContext(ctx)
    if err != nil {
        return err
    }
    var request StoreCreateRequest
    if err := ctx.BodyParser(&request); err != nil {
        return err
    }
    // 店舗作成
    output, err := c.storeService.Create(service.StoreServiceCommandInput{
        UserId: *userId,
        Name: request.Name,
    })
    if err != nil {
        return err
    }
    return ctx.JSON(c.toResponse(output))
}

// Update /* 店舗を更新
//	@Summary	店舗更新
//	@Tags		store
//	@Produce	json
//	@Param		storeId	path		string				true	"店舗ID"
//	@Param		request	body		StoreUpdateRequest	false	"店舗更新リクエスト"
//	@Success	200		{object}	StoreResponse
//	@Router		/stores/{storeId} [patch]
func (c StoreController) Update(ctx *fiber.Ctx) error {
    // JWTトークンからユーザーIDを取得
    userId, err := common.GetUserIdByContext(ctx)
    if err != nil {
        return err
    }
    // パスパラメータから店舗IDを取得
    id, err := uuid.Parse(ctx.Params("storeId"))
    if err != nil {
        return err
    }
    var request StoreUpdateRequest
    if err := ctx.BodyParser(&request); err != nil {
        return err
    }
    // 店舗作成
    output, err := c.storeService.Update(service.StoreServiceCommandUpdateInput{
        Id: id,
        UserId: *userId,
        Name: request.Name,
    })
    if err != nil {
        return err
    }
    return ctx.JSON(c.toResponse(output))
}

func (StoreController) toResponse(output *service.StoreServiceOutput) *StoreResponse {
    return &StoreResponse{
        Id: output.Id.String(),
        UserId: output.UserId.String(),
        Name: output.Name,
        CreatedAt: output.CreatedAt,
        UpdatedAt: output.UpdatedAt,
    }
}

type StoreRequest struct {
    Id uuid.UUID `json:"id"`
}

type StoreCreateRequest struct {
    Name string `json:"name"`
}

type StoreUpdateRequest struct {
    Name string `json:"name"`
}

type StoreListResponse struct {
    List []StoreResponse `json:"list"`
}

type StoreResponse struct {
    Id string `json:"id"`
    UserId string `json:"userId"`
    Name string `json:"name"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}
