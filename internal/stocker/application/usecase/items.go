package usecase

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"
	"time"

	"github.com/google/uuid"
)

type ItemUsecase struct {
	itemRepository repository.ItemRepository
}

// NewItemUsecase /* ItemUsecaseのプロバイダ
func NewItemUsecase(itemRepository repository.ItemRepository) ItemUsecase {
	return ItemUsecase{
		itemRepository,
	}
}

// GetItems /* アイテム一覧を取得する
func (s ItemUsecase) GetItems(storeId uuid.UUID, jancode, name *string) (*ItemUsecaseListOutput, error) {
	entities, err := s.itemRepository.Index(storeId, jancode, name)
	if err != nil {
		return nil, err
	}
	return s.toOutput(entities), nil
}

// CreateItem /* アイテムを作成する
func (s ItemUsecase) CreateItem(input ItemUsecaseInput) (*ItemUsecaseOutput, error) {
	entity, err := entity.NewItemEntity(input.StoreId, input.Name, input.JanCode)
	if err != nil {
		return nil, err
	}
	entity, err = s.itemRepository.Create(entity)
	if err != nil {
		return nil, err
	}
	output := ItemUsecaseOutput(*entity)
	return &output, nil
}

// UpdateItem /* アイテムを更新
func (s ItemUsecase) UpdateItem(input ItemUsecaseUpdateInput) (*ItemUsecaseOutput, error) {
	// 更新対象を取得
	entity, err := s.itemRepository.Select(input.StoreId, input.Id)
	if err != nil {
		return nil, err
	}
	// 更新
	entity.Update(input.Name, input.JanCode)
	entity, err = s.itemRepository.Update(entity)
	if err != nil {
		return nil, err
	}
	output := ItemUsecaseOutput(*entity)
	return &output, nil
}

func (ItemUsecase) toOutput(entities []entity.ItemEntity) *ItemUsecaseListOutput {
	var list []ItemUsecaseOutput
	for _, entity := range entities {
		list = append(
			list,
			ItemUsecaseOutput(entity),
		)
	}
	return &ItemUsecaseListOutput{
		List: list,
	}
}

type ItemUsecaseOutput struct {
	Id        uuid.UUID
	StoreId   uuid.UUID
	Name      string
	JanCode   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ItemUsecaseInput struct {
	StoreId uuid.UUID
	Name    string
	JanCode string
}

type ItemUsecaseUpdateInput struct {
	Id      uuid.UUID
	StoreId uuid.UUID
	Name    string
	JanCode string
}

type ItemUsecaseListOutput struct {
	List []ItemUsecaseOutput
}
