package service

import (
	"fmt"
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"
	"time"

	"github.com/google/uuid"
)

type StoreService struct {
    userRepository repository.UserRepository
	storeRepository repository.StoreRepository
}

func NewStoreService(userRepository repository.UserRepository, storeRepository repository.StoreRepository) StoreService {
    return StoreService{
        userRepository,
        storeRepository,
    }
}

func (s StoreService) Index(input StoreServiceQueryListInput) (*StoreServiceListOutput, error) {
    entities, err := s.storeRepository.Index(input.UserId)
    if err != nil {
        return nil, err
    }
    var list []StoreServiceOutput
    for _, entity := range(entities) {
        list = append(list, StoreServiceOutput(entity))
    }
    return &StoreServiceListOutput{
        List: list,
    }, nil
}

func (s StoreService) Select(input StoreServiceQueryInput) (*StoreServiceOutput, error) {
    entity, err := s.storeRepository.Select(input.Id)
    if err != nil {
        return nil, err
    }
    output := StoreServiceOutput(*entity)
    return &output, nil
}

func (s StoreService) Create(input StoreServiceCommandInput) (*StoreServiceOutput, error) {
    entity, err := entity.NewStoreEntity(input.UserId, input.Name)
    if err != nil {
        return nil, err
    }
    entity, err = s.storeRepository.Insert(entity)
    if err != nil {
        return nil, err
    }
    output := StoreServiceOutput(*entity)
    return &output, nil
}

func (s StoreService) Update(input StoreServiceCommandUpdateInput) (*StoreServiceOutput, error) {
    entity, err := s.storeRepository.Select(input.Id)
    if err != nil {
        return nil, err
    }
    if entity.UserId != input.UserId {
        return nil, fmt.Errorf("permission error")
    }
    entity.Update(input.Name)
    entity, err = s.storeRepository.Update(entity)
    if err != nil {
        return nil, err
    }
    output := StoreServiceOutput(*entity)
    return &output, nil
}

type StoreServiceQueryListInput struct {
    UserId uuid.UUID
}

type StoreServiceQueryInput struct {
    Id uuid.UUID
}

type StoreServiceCommandInput struct {
    UserId uuid.UUID
    Name string
}

type StoreServiceCommandUpdateInput struct {
    Id uuid.UUID
    UserId uuid.UUID
    Name string
}

type StoreServiceOutput struct {
    Id uuid.UUID
    UserId uuid.UUID
    Name string
    CreatedAt time.Time
    UpdatedAt time.Time
}

type StoreServiceListOutput struct {
    List []StoreServiceOutput
}