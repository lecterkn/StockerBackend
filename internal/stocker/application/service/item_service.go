package service

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"
	"time"

	"github.com/google/uuid"
)

type ItemService struct {
	ItemRepository repository.ItemRepository
}

func NewItemService(itemRepository repository.ItemRepository) ItemService {
	return ItemService{
		ItemRepository: itemRepository,
	}
}

func (s ItemService) GetItems() (*ItemServiceListOutput, error) {
	entities, err := s.ItemRepository.SelectItems()
	if err != nil {
		return nil, err
	}
	return s.toOutput(entities), nil
}

func (s ItemService) CreateItem(input ItemServiceInput) (*ItemServiceOutput, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	entity := &entity.ItemEntity{
		Id:        id,
		Name:      input.Name,
		JanCode:   input.JanCode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	entity, err = s.ItemRepository.Create(entity)
	if err != nil {
		return nil, err
	}
	output := ItemServiceOutput(*entity)
	return &output, nil
}

func (s ItemService) UpdateItem(input ItemServiceUpdateInput) (*ItemServiceOutput, error) {
	// 更新対象を取得
	entity, err := s.ItemRepository.Select(input.Id)
	if err != nil {
		return nil, err
	}
	// 更新
	entity.Update(input.Name, input.JanCode)
	entity, err = s.ItemRepository.Update(entity)
	output := ItemServiceOutput(*entity)
	return &output, nil
}

func (ItemService) toOutput(entities []entity.ItemEntity) *ItemServiceListOutput {
	var list []ItemServiceOutput
	for _, entity := range entities {
		list = append(
			list,
			ItemServiceOutput(entity),
		)
	}
	return &ItemServiceListOutput{
		List: list,
	}
}

type ItemServiceOutput struct {
	Id        uuid.UUID
	Name      string
	JanCode   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ItemServiceInput struct {
	Name    string
	JanCode string
}

type ItemServiceUpdateInput struct {
	Id      uuid.UUID
	Name    string
	JanCode string
}

type ItemServiceListOutput struct {
	List []ItemServiceOutput
}
