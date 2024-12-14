package usecase

import (
	"fmt"
	"h11/backend/internal/stocker/common"
	"h11/backend/internal/stocker/domain/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthorizationUsecase struct {
	userRepository repository.UserRepository
}

func NewAuthorizationUsecase(userRepository repository.UserRepository) AuthorizationUsecase {
	return AuthorizationUsecase{
		userRepository,
	}
}

func (s AuthorizationUsecase) Login(input AuthorizationUsecaseInput) (*AuthorizationUsecaseOutput, error) {
	// ユーザー取得
	userEntity, err := s.userRepository.SelectByName(input.Name)
	if err != nil {
		return nil, err
	}
	// パスワードチェック
	if !common.IsHashEquals(input.Password, userEntity.Password) {
		return nil, fmt.Errorf("password does not match")
	}
	jwtToken, err := common.GenerateJwt(jwt.MapClaims{
		"sub":  userEntity.Id,
		"name": userEntity.Name,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	if err != nil {
		return nil, err
	}
	return &AuthorizationUsecaseOutput{
		Token: *jwtToken,
	}, nil
}

type AuthorizationUsecaseInput struct {
	Name     string
	Password string
}

type AuthorizationUsecaseOutput struct {
	Token string
}
