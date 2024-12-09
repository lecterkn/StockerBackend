package implements

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/infrastructure/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StockInRepositoryImpl struct {
	database *gorm.DB
}

func NewStockInRepositoryImpl(database *gorm.DB) StockInRepositoryImpl {
	return StockInRepositoryImpl{
		database,
	}
}

// Create /* 入荷履歴を作成
func (r StockInRepositoryImpl) Create(entity *entity.StockInEntity) (*entity.StockInEntity, error) {
	model := r.toModel(entity)
	if err := r.database.Create(&model).Error; err != nil {
		return nil, err
	}
	entity, err := r.toEntity(model)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// Index /* 入庫履歴を取得
func (r StockInRepositoryImpl) Index(storeId uuid.UUID) ([]entity.StockInEntity, error) {
	var models []model.StockInModel
	if err := r.database.Where("store_id = ?", storeId[:]).Find(&models).Error; err != nil {
		return nil, err
	}
	var entities []entity.StockInEntity
	for _, model := range models {
		entity, err := r.toEntity(&model)
		if err != nil {
			continue
		}
		entities = append(entities, *entity)
	}
	return entities, nil
}

func (StockInRepositoryImpl) toEntity(model *model.StockInModel) (*entity.StockInEntity, error) {
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
	return &entity.StockInEntity{
		Id:        id,
		Place:     model.Place,
		StoreId:   storeId,
		ItemId:    itemId,
		Stocks:    model.Stocks,
		Price:     model.Price,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}

func (StockInRepositoryImpl) toModel(entity *entity.StockInEntity) *model.StockInModel {
	return &model.StockInModel{
		Id:        entity.Id[:],
		Place:     entity.Place,
		StoreId:   entity.StoreId[:],
		ItemId:    entity.ItemId[:],
		Stocks:    entity.Stocks,
		Price:     entity.Price,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
