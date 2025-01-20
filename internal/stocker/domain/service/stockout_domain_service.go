package service

import (
	"time"

	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"

	"github.com/google/uuid"
)

type StockOutDomainService struct {
	itemStockRepository repository.ItemStockRepository
	stockOutRepository  repository.StockOutRepository
}

func NewStockOutDomainService(itemStockRepository repository.ItemStockRepository, stockOutRepository repository.StockOutRepository) StockOutDomainService {
	return StockOutDomainService{
		itemStockRepository,
		stockOutRepository,
	}
}

func (s StockOutDomainService) AddStockOut(storeId uuid.UUID, itemEntity entity.ItemEntity, input StockOutDomainServiceCommandInput) (*StockOutDomainServiceCommandOutput, error) {
	stockOutEntity, err := entity.NewStockOutEntity(itemEntity, input.Price, input.Stocks)
	if err != nil {
		return nil, err
	}
	itemStockEntity, err := s.itemStockRepository.SelectByJancode(storeId, itemEntity.JanCode)
	if err != nil {
		return nil, err
	}
	itemStockEntity.Update(itemStockEntity.Price, itemStockEntity.Stock-input.Stocks, itemStockEntity.StockMin)
	if _, err := s.itemStockRepository.Update(itemStockEntity); err != nil {
		return nil, err
	}
	stockOutEntity, err = s.stockOutRepository.Create(stockOutEntity)
	if err != nil {
		return nil, err
	}
	output := s.toOutput(stockOutEntity)
	return &output, nil
}

func (StockOutDomainService) toOutput(entity *entity.StockOutEntity) StockOutDomainServiceCommandOutput {
	return StockOutDomainServiceCommandOutput{
		Id:        entity.Id,
		StoreId:   entity.Item.StoreId,
		ItemId:    entity.Item.Id,
		Price:     entity.Price,
		Stocks:    entity.Stocks,
		Name:      entity.Item.Name,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

type StockOutDomainServiceCommandOutput struct {
	Id        uuid.UUID
	StoreId   uuid.UUID
	ItemId    uuid.UUID
	Price     int
	Stocks    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StockOutDomainServiceCommandInput struct {
	Price  int
	Stocks int
}
