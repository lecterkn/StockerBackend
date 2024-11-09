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

// GetItems /* アイテム一覧を取得する
func (s ItemService) GetItems() (*ItemServiceListOutput, error) {
	entities, err := s.ItemRepository.SelectItems()
	if err != nil {
		return nil, err
	}
	return s.toOutput(entities), nil
}

// CreateItem /* アイテムを作成する
func (s ItemService) CreateItem(input ItemServiceInput) (*ItemServiceOutput, error) {
	entity, err := entity.NewItemEntity(input.Name, input.JanCode)
	if err != nil {
		return nil, err
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
