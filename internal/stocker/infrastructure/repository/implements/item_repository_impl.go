package implements

import (
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
        list = append(list, r.toEntity(model))
    }
    return list, nil
}

func (ItemRepositoryImpl) toEntity(model model.ItemModel) entity.ItemEntity {
    return entity.ItemEntity(model)
}
