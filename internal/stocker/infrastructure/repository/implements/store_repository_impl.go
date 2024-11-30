package implements

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/infrastructure/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StoreRepositoryImpl struct {
	database *gorm.DB
}

func NewStoreRepositoryImpl(database *gorm.DB) StoreRepositoryImpl {
    return StoreRepositoryImpl{
        database,
    }
}

func (r StoreRepositoryImpl) Index(userId uuid.UUID) ([]entity.StoreEntity, error) {
    var models []model.StoreModel
    if err := r.database.Where("userId = ?", userId[:]).Find(&models).Error; err != nil {
        return nil, err
    }
    var entities []entity.StoreEntity
    for _, model := range(models) {
        entity, err := r.toEntity(&model)
        if err != nil {
            return nil, err
        }
        entities = append(entities, *entity)
    }
    return entities, nil
}

func (r StoreRepositoryImpl) Select(id uuid.UUID) (*entity.StoreEntity, error) {
    var model *model.StoreModel
    if err := r.database.Where("id = ?", id[:]).First(model).Error; err != nil {
        return nil, err
    }
    entity, err := r.toEntity(model)
    if err != nil {
        return nil, err
    }
    return entity, nil
}

func (r StoreRepositoryImpl) Insert(entity *entity.StoreEntity) (*entity.StoreEntity, error) {
    model := r.toModel(entity)
    if err := r.database.Create(model).Error; err != nil {
        return nil, err
    }
    entity, err := r.toEntity(model)
    if err != nil {
        return nil, err
    }
    return entity, nil
}

func (r StoreRepositoryImpl) Update(entity *entity.StoreEntity) (*entity.StoreEntity, error) {
    model := r.toModel(entity)
    if err := r.database.Save(model).Error; err != nil {
        return nil, err
    }
    entity, err := r.toEntity(model)
    if err != nil {
        return nil, err
    }
    return entity, nil
}

func (StoreRepositoryImpl) toModel(entity *entity.StoreEntity) *model.StoreModel {
    return &model.StoreModel{
        Id: entity.Id[:],
        UserId: entity.UserId[:],
        Name: entity.Name,
        CreatedAt: entity.CreatedAt,
        UpdatedAt: entity.UpdatedAt,
    }
}

func (StoreRepositoryImpl) toEntity(model *model.StoreModel) (*entity.StoreEntity, error) {
    id, err := uuid.FromBytes(model.Id)
    if err != nil {
        return nil, err
    }
    userId, err := uuid.FromBytes(model.UserId)
    if err != nil {
        return nil, err
    }
    return &entity.StoreEntity{
        Id: id,
        UserId: userId,
        Name: model.Name,
        CreatedAt: model.CreatedAt,
        UpdatedAt: model.UpdatedAt,
    }, nil
}