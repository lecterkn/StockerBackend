package service

import (
	"time"

	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"

	"github.com/google/uuid"
)

type StockInDomainService struct {
	itemStockRepository repository.ItemStockRepository
	stockInRepository   repository.StockInRepository
}

func NewStockInDomainService(itemStockRepository repository.ItemStockRepository, stockInRepository repository.StockInRepository) StockInDomainService {
	return StockInDomainService{
		itemStockRepository,
		stockInRepository,
	}
}

func (s StockInDomainService) AddStockIn(storeId uuid.UUID, itemEntity entity.ItemEntity, input StockInDomainServiceCommandInput) (*StockInDomainServiceCommandOutput, error) {
	stockInEntity, err := entity.NewStockInEntity(itemEntity, input.Price, input.Stocks)
	if err != nil {
		return nil, err
	}
	itemStockEntity, err := s.itemStockRepository.SelectByJancode(storeId, itemEntity.JanCode)
	if err != nil {
		return nil, err
	}
	itemStockEntity.Update(itemStockEntity.Price, itemStockEntity.Stock+input.Stocks, itemStockEntity.StockMin)
	if _, err := s.itemStockRepository.Update(itemStockEntity); err != nil {
		return nil, err
	}
	stockInEntity, err = s.stockInRepository.Create(stockInEntity)
	if err != nil {
		return nil, err
	}
	output := s.toOutput(stockInEntity)
	return &output, nil
}

func (StockInDomainService) toOutput(entity *entity.StockInEntity) StockInDomainServiceCommandOutput {
	return StockInDomainServiceCommandOutput{
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

type StockInDomainServiceCommandOutput struct {
	Id        uuid.UUID
	StoreId   uuid.UUID
	ItemId    uuid.UUID
	Name      string
	Price     int
	Stocks    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type StockInDomainServiceCommandInput struct {
	Price  int
	Stocks int
}
