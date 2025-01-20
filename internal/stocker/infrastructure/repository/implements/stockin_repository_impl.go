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
	entity, err := r.toEntity(model, entity.Item)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// Index /* 入庫履歴を取得
func (r StockInRepositoryImpl) Index(storeId uuid.UUID) ([]entity.StockInEntity, error) {
	var models []model.StockInQueryModel
	if err := r.database.
		Model(&model.StockInModel{}).
		Select("stock_ins.*, items.name, items.jan_code, items.created_at as item_created_at, items.updated_at as item_updated_at").
		Joins("join items on stock_ins.item_id = items.id").
		Where("stock_ins.store_id = ?", storeId[:]).
		Order("stock_ins.id DESC").
		Find(&models).Error; err != nil {
		return nil, err
	}
	var entities []entity.StockInEntity
	for _, model := range models {
		entity, err := r.queryModelToEntity(&model)
		if err != nil {
			continue
		}
		entities = append(entities, *entity)
	}
	return entities, nil
}

func (StockInRepositoryImpl) queryModelToEntity(queryModel *model.StockInQueryModel) (*entity.StockInEntity, error) {
	id, err := uuid.FromBytes(queryModel.Id)
	if err != nil {
		return nil, err
	}
	storeId, err := uuid.FromBytes(queryModel.StoreId)
	if err != nil {
		return nil, err
	}
	itemId, err := uuid.FromBytes(queryModel.ItemId)
	if err != nil {
		return nil, err
	}
	return &entity.StockInEntity{
		Id:     id,
		Price:  queryModel.Price,
		Stocks: queryModel.Stocks,
		Item: entity.ItemEntity{
			Id:        itemId,
			StoreId:   storeId,
			Name:      queryModel.Name,
			JanCode:   queryModel.JanCode,
			CreatedAt: queryModel.ItemCreatedAt,
			UpdatedAt: queryModel.ItemUpdatedAt,
		},
		CreatedAt: queryModel.CreatedAt,
		UpdatedAt: queryModel.UpdatedAt,
	}, nil
}

func (StockInRepositoryImpl) toEntity(stockInModel *model.StockInModel, itemEntity entity.ItemEntity) (*entity.StockInEntity, error) {
	id, err := uuid.FromBytes(stockInModel.Id)
	if err != nil {
		return nil, err
	}
	return &entity.StockInEntity{
		Id:        id,
		Stocks:    stockInModel.Stocks,
		Price:     stockInModel.Price,
		Item:      itemEntity,
		CreatedAt: stockInModel.CreatedAt,
		UpdatedAt: stockInModel.UpdatedAt,
	}, nil
}

func (StockInRepositoryImpl) toModel(entity *entity.StockInEntity) *model.StockInModel {
	return &model.StockInModel{
		Id:        entity.Id[:],
		StoreId:   entity.Item.StoreId[:],
		ItemId:    entity.Item.Id[:],
		Stocks:    entity.Stocks,
		Price:     entity.Price,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
