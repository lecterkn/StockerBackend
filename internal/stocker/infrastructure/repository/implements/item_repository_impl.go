package implements

import (
    "github.com/google/uuid"
    "h11/backend/internal/stocker/domain/entity"
    "h11/backend/internal/stocker/infrastructure/model"

    "gorm.io/gorm"
)

type ItemRepositoryImpl struct {
    Database *gorm.DB
}

func NewItemRepositoryImpl(database *gorm.DB) ItemRepositoryImpl {
    return ItemRepositoryImpl{
        Database: database,
    }
}

// SelectItems /** アイテムを取得する
func (r ItemRepositoryImpl) SelectItems() ([]entity.ItemEntity, error) {
    var models []model.ItemModel
    err := r.Database.Find(&models).Error
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

func (r ItemRepositoryImpl) Select(id uuid.UUID) (*entity.ItemEntity, error) {
    var model model.ItemModel
    err := r.Database.Where("id = ?", id[:]).First(&model).Error
    if err != nil {
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
    err := r.Database.Create(&model).Error
    if err != nil {
        return nil, err
    }
    entity, err = r.toEntity(*model)
    if err != nil {
        return nil, err
    }
    return entity, nil
}

func (r ItemRepositoryImpl) Update(entity *entity.ItemEntity) (*entity.ItemEntity, error) {
    model := r.toModel(entity)
    err := r.Database.Save(&model).Error
    if err != nil {
        return nil, err
    }
    entity, err = r.toEntity(*model)
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
    return &entity.ItemEntity{
        Id:        id,
        Name:      model.Name,
        JanCode:   model.JanCode,
        CreatedAt: model.CreatedAt,
        UpdatedAt: model.UpdatedAt,
    }, nil
}

func (ItemRepositoryImpl) toModel(entity *entity.ItemEntity) *model.ItemModel {
    return &model.ItemModel{
        Id:        entity.Id[:],
        Name:      entity.Name,
        JanCode:   entity.JanCode,
        CreatedAt: entity.CreatedAt,
        UpdatedAt: entity.UpdatedAt,
    }
}
