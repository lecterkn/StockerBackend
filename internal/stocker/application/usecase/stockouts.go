package usecase

import (
	"time"

	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"

	"github.com/google/uuid"
)

type StockOutUsecase struct {
	itemRepository     repository.ItemRepository
	stockOutRepository repository.StockOutRepository
}

func NewStockOutUsecase(itemRepository repository.ItemRepository, stockOutRepository repository.StockOutRepository) StockOutUsecase {
	return StockOutUsecase{
		itemRepository,
		stockOutRepository,
	}
}

func (u StockOutUsecase) CreateStockOut(storeId uuid.UUID, input StockOutUsecaseCommandInput) (*StockOutUsecaseOutput, error) {
	itemEntity, err := u.itemRepository.Select(storeId, input.ItemId)
	if err != nil {
		return nil, err
	}
	entity, err := entity.NewStockOutEntity(*itemEntity, input.Price, input.Stocks)
	if err != nil {
		return nil, err
	}
	entity, err = u.stockOutRepository.Create(entity)
	if err != nil {
		return nil, err
	}
	output := u.toOutput(entity)
	return &output, nil
}

func (u StockOutUsecase) GetStockOuts(storeId uuid.UUID) (*StockOutUsecaseListOutput, error) {
	entities, err := u.stockOutRepository.Index(storeId)
	if err != nil {
		return nil, err
	}
	var list []StockOutUsecaseOutput
	for _, entity := range entities {
		output := u.toOutput(&entity)
		list = append(list, output)
	}
	return &StockOutUsecaseListOutput{
		list,
	}, nil
}

func (StockOutUsecase) toOutput(entity *entity.StockOutEntity) StockOutUsecaseOutput {
	return StockOutUsecaseOutput{
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

type StockOutUsecaseCommandInput struct {
	ItemId uuid.UUID
	Price  int
	Stocks int
}

type StockOutUsecaseListOutput struct {
	List []StockOutUsecaseOutput
}

type StockOutUsecaseOutput struct {
	Id        uuid.UUID
	StoreId   uuid.UUID
	ItemId    uuid.UUID
	Price     int
	Stocks    int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
