package service

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"
	"time"

	"github.com/google/uuid"
)

type ItemStockService struct {
	itemRepository repository.ItemStockRepository
}

// プロバイダ
func NewItemStockService(itemRepository repository.ItemStockRepository) ItemStockService {
	return ItemStockService{
		itemRepository: itemRepository,
	}
}

// Select /* 商品詳細を取得
func (s ItemStockService) Select(id uuid.UUID) (*ItemstockServiceOutput, error) {
	entity, err := s.itemRepository.Select(id)
	if err != nil {
		return nil, err
	}
	return &ItemstockServiceOutput{
		Id:        entity.ItemId,
		Place:     entity.Place,
		Stock:     entity.Stock,
		StockMin:  entity.StockMin,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}, nil
}

// Create /* 商品詳細を作成
func (s ItemStockService) Create(input ItemStockServiceInput) (*ItemstockServiceOutput, error) {
	entity, err := entity.NewItemStockEntity(input.ItemId, input.Place, input.Stock, input.StockMin)
	if err != nil {
		return nil, err
	}
	entity, err = s.itemRepository.Insert(entity)
	if err != nil {
		return nil, err
	}
	return &ItemstockServiceOutput{
		Id:        entity.ItemId,
		Place:     entity.Place,
		Stock:     entity.Stock,
		StockMin:  entity.StockMin,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}, nil
}

type ItemStockServiceInput struct {
	ItemId   uuid.UUID
	Place    string
	Stock    int
	StockMin int
}

type ItemstockServiceOutput struct {
	Id        uuid.UUID
	Place     string
	Stock     int
	StockMin  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
