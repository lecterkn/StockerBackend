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

// Select /* 商品詳細を取得
func (s ItemStockService) Select(id uuid.UUID) (*ItemStockServiceOutput, error) {
	entity, err := s.itemStockRepository.Select(id)
	if err != nil {
		return nil, err
	}
	output := ItemStockServiceOutput(*entity)
	return &output, nil
}

// Create /* 商品詳細を作成
func (s ItemStockService) Create(input ItemStockServiceInput) (*ItemStockServiceOutput, error) {
	entity, err := entity.NewItemStockEntity(input.ItemId, input.Place, input.Stock, input.StockMin)
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
	entity, err := s.itemStockRepository.Select(input.ItemId)
	if err != nil {
		return nil, err
	}
	// 商品詳細を更新
	entity = entity.Update(input.Place, input.Stock, input.StockMin)
	entity, err = s.itemStockRepository.Update(entity)
	if err != nil {
		return nil, err
	}
	output := ItemStockServiceOutput(*entity)
	return &output, nil
}

type ItemStockServiceInput struct {
	ItemId   uuid.UUID
	Place    string
	Stock    int
	StockMin int
}

type ItemStockServiceOutput struct {
	ItemId    uuid.UUID
	Place     string
	Stock     int
	StockMin  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
