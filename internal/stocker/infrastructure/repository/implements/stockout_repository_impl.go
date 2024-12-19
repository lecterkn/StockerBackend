package implements

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/infrastructure/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockOutRepositoryImpl struct {
	database *gorm.DB
}

func NewStockOutRepositoryImpl(database *gorm.DB) StockOutRepositoryImpl {
	return StockOutRepositoryImpl{
		database,
	}
}

func (r StockOutRepositoryImpl) Index(storeId uuid.UUID) ([]entity.StockOutEntity, error) {
	var models []model.StockOutQueryModel
	if err := r.database.Where("store_id = ?", storeId[:]).Find(&models).Error; err != nil {
		return nil, err
	}
	var entities []entity.StockOutEntity
	for _, model := range models {
		entity, err := r.fromQueryModelToEntity(&model)
		if err != nil {
			continue
		}
		entities = append(entities, *entity)
	}
	return entities, nil
}

func (r StockOutRepositoryImpl) Create(entity *entity.StockOutEntity) (*entity.StockOutEntity, error) {
	model := r.toModel(entity)
	if err := r.database.Create(&model).Error; err != nil {
		return nil, err
	}
	entity, err := r.toEntity(model, entity.Item)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (StockOutRepositoryImpl) fromQueryModelToEntity(model *model.StockOutQueryModel) (*entity.StockOutEntity, error) {
	id, err := uuid.FromBytes(model.Id)
	if err != nil {
		return nil, err
	}
	storeId, err := uuid.FromBytes(model.StoreId)
	if err != nil {
		return nil, err
	}
	itemId, err := uuid.FromBytes(model.ItemId)
	if err != nil {
		return nil, err
	}
	return &entity.StockOutEntity{
		Id:        id,
		Place:     model.Place,
		Price:     model.Price,
		Stocks:    model.Stocks,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
		Item: entity.ItemEntity{
			Id:        itemId,
			StoreId:   storeId,
			Name:      model.Name,
			JanCode:   model.JanCode,
			CreatedAt: model.ItemCreatedAt,
			UpdatedAt: model.ItemUpdatedAt,
		},
	}, nil
}

func (StockOutRepositoryImpl) toEntity(model *model.StockOutModel, itemEntity entity.ItemEntity) (*entity.StockOutEntity, error) {
	id, err := uuid.FromBytes(model.Id)
	if err != nil {
		return nil, err
	}
	return &entity.StockOutEntity{
		Id:        id,
		Place:     model.Place,
		Price:     model.Price,
		Stocks:    model.Stocks,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}

func (StockOutRepositoryImpl) toModel(entity *entity.StockOutEntity) *model.StockOutModel {
	return &model.StockOutModel{
		Id:        entity.Id[:],
		StoreId:   entity.Item.StoreId[:],
		ItemId:    entity.Item.Id[:],
		Place:     entity.Place,
		Price:     entity.Price,
		Stocks:    entity.Stocks,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
