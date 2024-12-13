package usecase

import (
	"time"

	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"
	"h11/backend/internal/stocker/domain/service"

	"github.com/google/uuid"
)

type ItemStockUsecase struct {
	itemStockDomainUsecase service.ItemStockDomainService
	itemStockRepository    repository.ItemStockRepository
	itemRepository         repository.ItemRepository
}

// プロバイダ
func NewItemStockUsecase(itemStockRepository repository.ItemStockRepository, itemRepository repository.ItemRepository, itemStockDomainUsecase service.ItemStockDomainService) ItemStockUsecase {
	return ItemStockUsecase{
		itemStockDomainUsecase,
		itemStockRepository,
		itemRepository,
	}
}

// Index /* 商品詳細一覧を取得する
func (s ItemStockUsecase) Index(input ItemStockUsecaseQueryListInput) (*ItemStockUsecaseListOutput, error) {
	entities, err := s.itemStockRepository.Index(input.StoreId)
	if err != nil {
		return nil, err
	}
	var list []ItemStockUsecaseOutput
	for _, entity := range entities {
		list = append(list, *s.toOutput(&entity))
	}
	return &ItemStockUsecaseListOutput{
		list,
	}, nil
}

// Select /* 商品詳細を取得
func (s ItemStockUsecase) Select(input ItemStockUsecaseQueryInput) (*ItemStockUsecaseOutput, error) {
	entity, err := s.itemStockRepository.Select(input.StoreId, input.ItemId)
	if err != nil {
		return nil, err
	}
	return s.toOutput(entity), nil
}

// Create /* 商品詳細を作成
func (s ItemStockUsecase) Create(storeId uuid.UUID, input ItemStockUsecaseInput) (*ItemStockUsecaseOutput, error) {
	entity, err := entity.NewItemStockEntity(storeId, input.Name, input.JanCode, input.Price, input.Stock, input.StockMin)
	if err != nil {
		return nil, err
	}
	entity, err = s.itemStockDomainUsecase.CreateItemStock(entity)
	if err != nil {
		return nil, err
	}
	return s.toOutput(entity), nil
}

func (s ItemStockUsecase) Update(storeId, itemId uuid.UUID, input ItemStockUsecaseUpdateInput) (*ItemStockUsecaseOutput, error) {
	// 商品詳細を取得
	entity, err := s.itemStockRepository.Select(storeId, itemId)
	if err != nil {
		return nil, err
	}
	// 商品詳細を更新
	entity.Update(input.Price, input.Stock, input.StockMin)
	entity, err = s.itemStockRepository.Update(entity)
	if err != nil {
		return nil, err
	}
	return s.toOutput(entity), nil
}

func (ItemStockUsecase) toOutput(entity *entity.ItemStockEntity) *ItemStockUsecaseOutput {
	return &ItemStockUsecaseOutput{
		Name:      entity.Item.Name,
		JanCode:   entity.Item.JanCode,
		ItemId:    entity.Item.Id,
		StoreId:   entity.Item.StoreId,
		Stock:     entity.Stock,
		StockMin:  entity.StockMin,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

type ItemStockUsecaseQueryListInput struct {
	StoreId uuid.UUID
}

type ItemStockUsecaseQueryInput struct {
	StoreId uuid.UUID
	ItemId  uuid.UUID
}

type ItemStockUsecaseInput struct {
	Name     string
	JanCode  string
	Price    *int
	Stock    int
	StockMin *int
}

type ItemStockUsecaseUpdateInput struct {
	Price    *int
	Stock    int
	StockMin *int
}

type ItemStockUsecaseListOutput struct {
	List []ItemStockUsecaseOutput
}

type ItemStockUsecaseOutput struct {
	Name      string
	JanCode   string
	ItemId    uuid.UUID
	StoreId   uuid.UUID
	Price     *int
	Stock     int
	StockMin  *int
	CreatedAt time.Time
	UpdatedAt time.Time
}
