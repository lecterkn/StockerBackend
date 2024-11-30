package controller

import (
	"h11/backend/internal/stocker/application/service"
	"time"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService service.UserService
    authorizationService service.AuthorizationService
}

func NewUserController(userService service.UserService, authorizationService service.AuthorizationService) UserController {
    return UserController{
        userService,
        authorizationService,
    }
}

// Create /* ユーザーを新規作成
//	@Summary	ユーザー作成
//	@Tags		user
//	@Produce	json
//	@Param		request	body		UserRequest		false	"ユーザー作成リクエスト"
//	@Success	200		{object}	UserResponse	
//	@Router		/register [post]
func (c UserController) Create(ctx *fiber.Ctx) error {
    var request UserRequest
    if err := ctx.BodyParser(&request); err != nil {
        return err
    }
    output, err := c.userService.CreateUser(service.UserServiceInput{
        Name: request.Name,
        Password: request.Password,
    })
    if err != nil {
        return err
    }
    return ctx.JSON(UserResponse(*c.toResponse(output)))
}

// Login /* ユーザーにログイン
//	@Summary	ユーザーログイン
//	@Tags		user
//	@Produce	json
//	@Param		request	body		UserRequest			false	"ユーザーログインリクエスト"
//	@Success	200		{object}	UserLoginResponse	
//	@Router		/login [post]
func (c UserController) Login(ctx *fiber.Ctx) error {
    var request UserRequest
    if err := ctx.BodyParser(&request); err != nil {
        return err
    }
    output, err := c.authorizationService.Login(service.AuthorizationServiceInput(request))
    if err != nil {
        return err
    }
    return ctx.JSON(UserLoginResponse(*output))
}

func (UserController) toResponse(output *service.UserServiceOutput) *UserResponse {
    return &UserResponse{
        Id: output.Id.String(),
        Name: output.Name,
        Password: string(output.Password),
        CreateAt: output.CreatedAt,
        UpdatedAt: output.UpdatedAt,
    }
}

type UserRequest struct {
    Name string `json:"name"`
    Password string `json:"password"`
}

type UserResponse struct {
    Id string
    Name string
    Password string `json:"-"`
    CreateAt time.Time
    UpdatedAt time.Time
}

type UserLoginResponse struct {
    Token string
}