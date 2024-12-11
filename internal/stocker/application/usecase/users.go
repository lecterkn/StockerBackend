package usecase

import (
	"time"

	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"

	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) UserUsecase {
	return UserUsecase{
		userRepository,
	}
}

func (s UserUsecase) GetUserByName(name string) (*UserUsecaseOutput, error) {
	entity, err := s.userRepository.SelectByName(name)
	if err != nil {
		return nil, err
	}
	output := UserUsecaseOutput(*entity)
	return &output, nil
}

func (s UserUsecase) CreateUser(input UserUsecaseInput) (*UserUsecaseOutput, error) {
	// ユーザーエンティティ作成
	entity, err := entity.NewUserEntity(input.Name, input.Password)
	if err != nil {
		return nil, err
	}
	// ユーザーを挿入
	entity, err = s.userRepository.Insert(entity)
	if err != nil {
		return nil, err
	}
	// アウトプット
	output := UserUsecaseOutput(*entity)
	return &output, nil
}

type UserUsecaseInput struct {
	Name     string
	Password string
}

type UserUsecaseOutput struct {
	Id        uuid.UUID
	Name      string
	Password  []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}
