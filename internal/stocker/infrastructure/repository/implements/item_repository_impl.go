package implements

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/infrastructure/model"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type ItemRepositoryImpl struct {
	database *gorm.DB
}

// NewItemRepositoryImpl /* ItemRepositoryImplのプロバイダ
func NewItemRepositoryImpl(database *gorm.DB) ItemRepositoryImpl {
	return ItemRepositoryImpl{
		database,
	}
}

// SelectItems /** アイテムを取得する
func (r ItemRepositoryImpl) Index(storeId uuid.UUID, jancode, name *string) ([]entity.ItemEntity, error) {
	models, err := r.finds(storeId, jancode, name)
	if err != nil {
		return nil, err
	}
	var list []entity.ItemEntity
	for _, model := range models {
		entity, err := r.toEntity(model)
		if err != nil {
			return nil, err
		}
		list = append(list, *entity)
	}
	return list, nil
}

func (r ItemRepositoryImpl) finds(storeId uuid.UUID, jancode, name *string) ([]model.ItemModel, error) {
	var models []model.ItemModel
	r.database.Where("store_id = ?", storeId[:])
	if jancode != nil {
		r.database.Where("jan_code = ?", *jancode)
	}
	if name != nil {
		r.database.Where("name = ?", *name)
	}
	if err := r.database.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}

// Select /* idからアイテムを取得する
func (r ItemRepositoryImpl) Select(storeId, id uuid.UUID) (*entity.ItemEntity, error) {
	var model model.ItemModel
	if err := r.database.Where("store_id = ? AND id = ?", storeId[:], id[:]).First(&model).Error; err != nil {
		return nil, err
	}
	entity, err := r.toEntity(model)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// Create /** アイテムを作成する
func (r ItemRepositoryImpl) Create(entity *entity.ItemEntity) (*entity.ItemEntity, error) {
	model := r.toModel(entity)
	if err := r.database.Create(&model).Error; err != nil {
		return nil, err
	}
	entity, err := r.toEntity(*model)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

// Update /* アイテムを更新する
func (r ItemRepositoryImpl) Update(entity *entity.ItemEntity) (*entity.ItemEntity, error) {
	model := r.toModel(entity)
	if err := r.database.Save(&model).Error; err != nil {
		return nil, err
	}
	entity, err := r.toEntity(*model)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (ItemRepositoryImpl) toEntity(model model.ItemModel) (*entity.ItemEntity, error) {
	id, err := uuid.FromBytes(model.Id)
	if err != nil {
		return nil, err
	}
	storeId, err := uuid.FromBytes(model.StoreId)
	if err != nil {
		return nil, err
	}
	return &entity.ItemEntity{
		Id:        id,
		StoreId:   storeId,
		Name:      model.Name,
		JanCode:   model.JanCode,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}

func (ItemRepositoryImpl) toModel(entity *entity.ItemEntity) *model.ItemModel {
	return &model.ItemModel{
		Id:        entity.Id[:],
		StoreId:   entity.StoreId[:],
		Name:      entity.Name,
		JanCode:   entity.JanCode,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
