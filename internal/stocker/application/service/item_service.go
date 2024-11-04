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

type ItemServiceListOutput struct {
	list []ItemServiceOutput
}

func (s ItemService) GetItems() (*ItemServiceListOutput, error) {
	entities, err := s.ItemRepository.SelectItems()
	if err != nil {
		return nil, err
	}
	return s.toOutput(entities), nil
}

func (s ItemService) CreateItem(input ItemServiceInput) *ItemServiceOutput {
	id, err := uuid.NewV7()
	if err != nil {
		return nil
	}
	entity := &entity.ItemEntity{
		Id:        id,
		Name:      input.Name,
		JanCode:   input.JanCode,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	entity, err = s.ItemRepository.Create(*entity)
	if err != nil {
		return nil
	}
	output := ItemServiceOutput(*entity)
	return &output
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
		list: list,
	}
}
