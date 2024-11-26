package implements

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/infrastructure/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemStockRepositoryImpl struct {
	database *gorm.DB
}

// NewItemStockRepositoryImpl /* プロバイダ
func NewItemStockRepositoryImpl(database *gorm.DB) ItemStockRepositoryImpl {
	return ItemStockRepositoryImpl{
		database,
	}
}

// Select /* IDから商品を取得
func (r ItemStockRepositoryImpl) Select(id uuid.UUID) (*entity.ItemStockEntity, error) {
	var model model.ItemStockModel
	err := r.database.Where("id = ?", id[:]).Select(&model).Error
	if err != nil {
		return nil, err
	}
	entity, err := r.ToEntity(&model)
	if err != nil {
		return nil, err
	}
	return entity, err
}

// Insert /* 商品をデータベースに挿入
func (r ItemStockRepositoryImpl) Insert(entity *entity.ItemStockEntity) (*entity.ItemStockEntity, error) {
	model := r.ToModel(entity)
	err := r.database.Create(model).Error
	if err != nil {
		return nil, err
	}
	entity, err = r.ToEntity(model)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r ItemStockRepositoryImpl) Update(entity *entity.ItemStockEntity) (*entity.ItemStockEntity, error) {
	model := r.ToModel(entity)
	err := r.database.Save(model).Error
	if err != nil {
		return nil, err
	}
	entity, err = r.ToEntity(model)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (ItemStockRepositoryImpl) ToModel(entity *entity.ItemStockEntity) *model.ItemStockModel {
	return &model.ItemStockModel{
		ItemId:    entity.ItemId[:],
		Place:     entity.Place,
		Stock:     entity.Stock,
		StockMin:  entity.StockMin,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
func (ItemStockRepositoryImpl) ToEntity(model *model.ItemStockModel) (*entity.ItemStockEntity, error) {
	id, err := uuid.ParseBytes(model.ItemId)
	if err != nil {
		return nil, err
	}
	return &entity.ItemStockEntity{
		ItemId:    id,
		Place:     model.Place,
		Stock:     model.Stock,
		StockMin:  model.StockMin,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}
