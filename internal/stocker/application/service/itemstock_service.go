package service

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"
	"time"

	"github.com/google/uuid"
)

type ItemStockService struct {
	itemStockRepository repository.ItemStockRepository
}

// プロバイダ
func NewItemStockService(itemStockRepository repository.ItemStockRepository) ItemStockService {
	return ItemStockService{
		itemStockRepository,
	}
}

// Index /* 商品詳細一覧を取得する
func (s ItemStockService) Index(input ItemStockServiceQueryListInput) (*ItemStockServiceListOutput, error) {
	entities, err := s.itemStockRepository.Index(input.StoreId)
	if err != nil {
		return nil, err
	}
	var list []ItemStockServiceOutput
	for _, entity := range entities {
		list = append(list, ItemStockServiceOutput(entity))
	}
	return &ItemStockServiceListOutput{
		list,
	}, nil
}

// Select /* 商品詳細を取得
func (s ItemStockService) Select(input ItemStockServiceQueryInput) (*ItemStockServiceOutput, error) {
	entity, err := s.itemStockRepository.Select(input.StoreId, input.Id)
	if err != nil {
		return nil, err
	}
	output := ItemStockServiceOutput(*entity)
	return &output, nil
}

// Create /* 商品詳細を作成
func (s ItemStockService) Create(input ItemStockServiceInput) (*ItemStockServiceOutput, error) {
	entity, err := entity.NewItemStockEntity(input.ItemId, input.StoreId, input.Place, input.Price, input.Stock, input.StockMin)
	if err != nil {
		return nil, err
	}
	entity, err = s.itemStockRepository.Insert(entity)
	if err != nil {
		return nil, err
	}
	output := ItemStockServiceOutput(*entity)
	return &output, nil
}

func (s ItemStockService) Update(input ItemStockServiceInput) (*ItemStockServiceOutput, error) {
	// 商品詳細を取得
	entity, err := s.itemStockRepository.Select(input.StoreId, input.ItemId)
	if err != nil {
		return nil, err
	}
	// 商品詳細を更新
	entity.Update(input.Place, input.Price, input.Stock, input.StockMin)
	entity, err = s.itemStockRepository.Update(entity)
	if err != nil {
		return nil, err
	}
	output := ItemStockServiceOutput(*entity)
	return &output, nil
}

type ItemStockServiceQueryListInput struct {
	StoreId uuid.UUID
}

type ItemStockServiceQueryInput struct {
	StoreId uuid.UUID
	Id      uuid.UUID
}

type ItemStockServiceInput struct {
	StoreId  uuid.UUID
	ItemId   uuid.UUID
	Place    string
	Price    *int
	Stock    int
	StockMin int
}

type ItemStockServiceListOutput struct {
	List []ItemStockServiceOutput
}

type ItemStockServiceOutput struct {
	ItemId    uuid.UUID
	StoreId   uuid.UUID
	Place     string
	Price     *int
	Stock     int
	StockMin  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
