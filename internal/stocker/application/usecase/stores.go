package usecase

import (
	"fmt"
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"
	"time"

	"github.com/google/uuid"
)

type StoreUsecase struct {
	userRepository  repository.UserRepository
	storeRepository repository.StoreRepository
}

func NewStoreUsecase(userRepository repository.UserRepository, storeRepository repository.StoreRepository) StoreUsecase {
	return StoreUsecase{
		userRepository,
		storeRepository,
	}
}

func (s StoreUsecase) Index(input StoreUsecaseQueryListInput) (*StoreUsecaseListOutput, error) {
	entities, err := s.storeRepository.Index(input.UserId)
	if err != nil {
		return nil, err
	}
	var list []StoreUsecaseOutput
	for _, entity := range entities {
		list = append(list, StoreUsecaseOutput(entity))
	}
	return &StoreUsecaseListOutput{
		List: list,
	}, nil
}

func (s StoreUsecase) Select(input StoreUsecaseQueryInput) (*StoreUsecaseOutput, error) {
	entity, err := s.storeRepository.Select(input.Id)
	if err != nil {
		return nil, err
	}
	output := StoreUsecaseOutput(*entity)
	return &output, nil
}

func (s StoreUsecase) Create(input StoreUsecaseCommandInput) (*StoreUsecaseOutput, error) {
	entity, err := entity.NewStoreEntity(input.UserId, input.Name)
	if err != nil {
		return nil, err
	}
	entity, err = s.storeRepository.Insert(entity)
	if err != nil {
		return nil, err
	}
	output := StoreUsecaseOutput(*entity)
	return &output, nil
}

func (s StoreUsecase) Update(input StoreUsecaseCommandUpdateInput) (*StoreUsecaseOutput, error) {
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
	output := StoreUsecaseOutput(*entity)
	return &output, nil
}

type StoreUsecaseQueryListInput struct {
	UserId uuid.UUID
}

type StoreUsecaseQueryInput struct {
	Id uuid.UUID
}

type StoreUsecaseCommandInput struct {
	UserId uuid.UUID
	Name   string
}

type StoreUsecaseCommandUpdateInput struct {
	Id     uuid.UUID
	UserId uuid.UUID
	Name   string
}

type StoreUsecaseOutput struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StoreUsecaseListOutput struct {
	List []StoreUsecaseOutput
}
