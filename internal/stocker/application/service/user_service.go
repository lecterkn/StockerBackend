package service

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
    return UserService{
        userRepository,
    }
}

func (s UserService) GetUserByName(name string) (*UserServiceOutput, error) {
    entity, err := s.userRepository.SelectByName(name)
    if err != nil {
        return nil, err
    }
    output := UserServiceOutput(*entity)
    return &output, nil
}

func (s UserService) CreateUser(input UserServiceInput) (*UserServiceOutput, error) {
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
    output := UserServiceOutput(*entity)
    return &output, nil
}

type UserServiceInput struct {
    Name string
    Password string
}

type UserServiceOutput struct {
    Id uuid.UUID
    Name string
    Password []byte
    CreatedAt time.Time
    UpdatedAt time.Time
}