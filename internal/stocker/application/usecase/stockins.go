package usecase

import (
	"time"

	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"
	"h11/backend/internal/stocker/domain/service"

	"github.com/google/uuid"
)

type StockInUsecase struct {
	stockInDomainService service.StockInDomainService
	stockInRepository    repository.StockInRepository
	itemRepository       repository.ItemRepository
}

func NewStockInUsecase(stockInDomainService service.StockInDomainService, stockInRepository repository.StockInRepository, itemRepository repository.ItemRepository) StockInUsecase {
	return StockInUsecase{
		stockInDomainService,
		stockInRepository,
		itemRepository,
	}
}

// GetStockIns /* 店舗IDから入庫履歴を取得
func (s StockInUsecase) GetStockIns(storeId uuid.UUID) (*StockInListOutput, error) {
	entities, err := s.stockInRepository.Index(storeId)
	if err != nil {
		return nil, err
	}
	var list []StockInOutput
	for _, entity := range entities {
		list = append(list, *s.toOutput(&entity))
	}
	output := StockInListOutput{
		list,
	}
	return &output, nil
}

// CreateStockIn /* 入庫履歴を作成
func (s StockInUsecase) CreateStockIn(storeId uuid.UUID, input StockInCommandInput) (*StockInOutput, error) {
	itemEntity, err := s.itemRepository.Select(storeId, input.ItemId)
	if err != nil {
		return nil, err
	}
	stockInDomainServiceOutput, err := s.stockInDomainService.AddStockIn(storeId, *itemEntity, service.StockInDomainServiceCommandInput{
		Price:  input.Price,
		Stocks: input.Stocks,
	})
	if err != nil {
		return nil, err
	}
	output := StockInOutput(*stockInDomainServiceOutput)
	return &output, nil
}

func (StockInUsecase) toOutput(entity *entity.StockInEntity) *StockInOutput {
	return &StockInOutput{
		Id:        entity.Id,
		StoreId:   entity.Item.StoreId,
		ItemId:    entity.Item.Id,
		Name:      entity.Item.Name,
		Price:     entity.Price,
		Stocks:    entity.Stocks,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

type StockInCommandInput struct {
	ItemId uuid.UUID
	Price  int
	Stocks int
}

type StockInListOutput struct {
	List []StockInOutput
}

type StockInOutput struct {
	Id        uuid.UUID
	StoreId   uuid.UUID
	ItemId    uuid.UUID
	Name      string
	Price     int
	Stocks    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
