package service

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"
	"time"

	"github.com/google/uuid"
)

type StockInService struct {
	stockInRepository repository.StockInRepository
}

func NewStockInService(stockInRepository repository.StockInRepository) StockInService {
	return StockInService{
		stockInRepository,
	}
}

// GetStockIns /* 店舗IDから入庫履歴を取得
func (s StockInService) GetStockIns(storeId uuid.UUID) (*StockInListOutput, error) {
	entities, err := s.stockInRepository.Index(storeId)
	if err != nil {
		return nil, err
	}
	var list []StockInOutput
	for _, entity := range entities {
		list = append(list, StockInOutput(entity))
	}
	output := StockInListOutput{
		list,
	}
	return &output, nil
}

// CreateStockIn /* 入庫履歴を作成
func (s StockInService) CreateStockIn(storeId uuid.UUID, input StockInCommandInput) (*StockInOutput, error) {
	entity, err := entity.NewStockInEntity(input.Place, storeId, input.ItemId, input.Price, input.Stocks)
	if err != nil {
		return nil, err
	}
	entity, err = s.stockInRepository.Create(entity)
	if err != nil {
		return nil, err
	}
	output := StockInOutput(*entity)
	return &output, nil
}

type StockInCommandInput struct {
	Place  string
	ItemId uuid.UUID
	Price  int
	Stocks int
}

type StockInListOutput struct {
	List []StockInOutput
}

type StockInOutput struct {
	Id        uuid.UUID
	Place     string
	StoreId   uuid.UUID
	ItemId    uuid.UUID
	Price     int
	Stocks    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
