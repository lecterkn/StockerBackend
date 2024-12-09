package service

import (
	"fmt"
	"h11/backend/internal/stocker/common"
	"h11/backend/internal/stocker/domain/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthorizationService struct {
	userRepository repository.UserRepository
}

func NewAuthorizationService(userRepository repository.UserRepository) AuthorizationService {
	return AuthorizationService{
		userRepository,
	}
}

func (s AuthorizationService) Login(input AuthorizationServiceInput) (*AuthorizationServiceOutput, error) {
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
	return &AuthorizationServiceOutput{
		Token: *jwtToken,
	}, nil
}

type AuthorizationServiceInput struct {
	Name     string
	Password string
}

type AuthorizationServiceOutput struct {
	Token string
}
