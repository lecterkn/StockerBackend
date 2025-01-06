package controller

import (
	"time"

	"h11/backend/internal/stocker/application/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userUsecase          usecase.UserUsecase
	authorizationUsecase usecase.AuthorizationUsecase
}

func NewUserController(userUsecase usecase.UserUsecase, authorizationUsecase usecase.AuthorizationUsecase) UserController {
	return UserController{
		userUsecase,
		authorizationUsecase,
	}
}

// Create /* ユーザーを新規作成
//
//	@Summary	ユーザー作成
//	@Tags		user
//	@Produce	json
//	@Param		request	body		UserRequest	false	"ユーザー作成リクエスト"
//	@Success	200		{object}	UserResponse
//	@Router		/register [post]
func (c UserController) Create(ctx *fiber.Ctx) error {
	var request UserRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}
	output, err := c.userUsecase.CreateUser(usecase.UserUsecaseInput{
		Name:     request.Name,
		Password: request.Password,
	})
	if err != nil {
		return err
	}
	return ctx.JSON(UserResponse(*c.toResponse(output)))
}

// Login /* ユーザーにログイン
//
//	@Summary	ユーザーログイン
//	@Tags		user
//	@Produce	json
//	@Param		request	body		UserRequest	false	"ユーザーログインリクエスト"
//	@Success	200		{object}	UserLoginResponse
//	@Router		/login [post]
func (c UserController) Login(ctx *fiber.Ctx) error {
	var request UserRequest
	if err := ctx.BodyParser(&request); err != nil {
		return err
	}
	output, err := c.authorizationUsecase.Login(usecase.AuthorizationUsecaseInput(request))
	if err != nil {
		return err
	}
	return ctx.JSON(UserLoginResponse(*output))
}

func (UserController) toResponse(output *usecase.UserUsecaseOutput) *UserResponse {
	return &UserResponse{
		Id:        output.Id.String(),
		Name:      output.Name,
		Password:  string(output.Password),
		CreateAt:  output.CreatedAt,
		UpdatedAt: output.UpdatedAt,
	}
}

type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	Id        string    `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Password  string    `json:"-" validate:"required"`
	CreateAt  time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

type UserLoginResponse struct {
	Token string `json:"token" validate:"required"`
}
