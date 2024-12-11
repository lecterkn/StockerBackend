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

// Index /* 商品詳細一覧を取得
func (r ItemStockRepositoryImpl) Index(storeId uuid.UUID) ([]entity.ItemStockEntity, error) {
	var models []model.ItemStockQueryModel
	if err := r.database.
		Model(&model.ItemStockModel{}).
		Select("item_stocks.*, items.name, items.jan_code, items.created_at as item_created_at, items.updated_at as item_updated_at").
		Joins("JOIN items ON items.id = item_stocks.item_id").
		Where("items.store_id = ?", storeId[:]).
		Find(&models).Error; err != nil {
		return nil, err
	}
	var entities []entity.ItemStockEntity
	for _, model := range models {
		entity, err := r.queryModelToEntity(&model)
		if err != nil {
			return nil, err
		}
		entities = append(entities, *entity)
	}
	return entities, nil
}

// Select /* IDから商品を取得
func (r ItemStockRepositoryImpl) Select(storeId, id uuid.UUID) (*entity.ItemStockEntity, error) {
	var itemStockQueryModel model.ItemStockQueryModel
	if err := r.database.
		Model(&model.ItemStockModel{}).
		Select("item_stocks.*, items.name, items.jan_code, items.created_at as item_created_at, items.updated_at as item_updated_at").
		Where("store_id = ? AND id = ?", storeId[:], id[:]).
		Joins("JOIN items ON items.id = item_stocks.item_id").
		First(&itemStockQueryModel).Error; err != nil {
		return nil, err
	}
	entity, err := r.queryModelToEntity(&itemStockQueryModel)
	if err != nil {
		return nil, err
	}
	return entity, err
}

// Insert /* 商品をデータベースに挿入
func (r ItemStockRepositoryImpl) Insert(entity *entity.ItemStockEntity) (*entity.ItemStockEntity, error) {
	itemEntity := r.item
	model := r.toModel(entity)
	if err := r.database.Create(model).Error; err != nil {
		return nil, err
	}
	entity, err := r.toEntity(model, entity.Item)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r ItemStockRepositoryImpl) Update(entity *entity.ItemStockEntity) (*entity.ItemStockEntity, error) {
	model := r.toModel(entity)
	if err := r.database.Save(model).Error; err != nil {
		return nil, err
	}
	entity, err := r.toEntity(model, entity.Item)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (ItemStockRepositoryImpl) toModel(entity *entity.ItemStockEntity) *model.ItemStockModel {
	return &model.ItemStockModel{
		ItemId:    entity.Item.Id[:],
		StoreId:   entity.Item.StoreId[:],
		Place:     entity.Place,
		Price:     entity.Price,
		Stock:     entity.Stock,
		StockMin:  entity.StockMin,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func (ItemStockRepositoryImpl) toEntity(model *model.ItemStockModel, itemEntity entity.ItemEntity) (*entity.ItemStockEntity, error) {
	return &entity.ItemStockEntity{
		Item:      itemEntity,
		Place:     model.Place,
		Price:     model.Price,
		Stock:     model.Stock,
		StockMin:  model.StockMin,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}

func (ItemStockRepositoryImpl) queryModelToEntity(model *model.ItemStockQueryModel) (*entity.ItemStockEntity, error) {
	id, err := uuid.FromBytes(model.ItemId)
	if err != nil {
		return nil, err
	}
	storeId, err := uuid.FromBytes(model.StoreId)
	if err != nil {
		return nil, err
	}
	return &entity.ItemStockEntity{
		Item: entity.ItemEntity{
			Id:        id,
			StoreId:   storeId,
			Name:      model.Name,
			JanCode:   model.JanCode,
			CreatedAt: model.ItemCreatedAt,
			UpdatedAt: model.ItemUpdatedAt,
		},
		Place:     model.Place,
		Price:     model.Price,
		Stock:     model.Stock,
		StockMin:  model.StockMin,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}
