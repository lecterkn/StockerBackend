package implements

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/infrastructure/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	database *gorm.DB
}

func NewUserRepositoryImpl(database *gorm.DB) UserRepositoryImpl {
	return UserRepositoryImpl{
		database,
	}
}

func (r UserRepositoryImpl) Select(id uuid.UUID) (*entity.UserEntity, error) {
	var model model.UserModel
	if err := r.database.Where("id = ?", id[:]).First(&model).Error; err != nil {
		return nil, err
	}
	entity, err := r.toEntity(&model)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r UserRepositoryImpl) SelectByName(name string) (*entity.UserEntity, error) {
	var model model.UserModel
	if err := r.database.Where("name = ?", name).First(&model).Error; err != nil {
		return nil, err
	}
	entity, err := r.toEntity(&model)
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r UserRepositoryImpl) Insert(entity *entity.UserEntity) (*entity.UserEntity, error) {
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

func (r UserRepositoryImpl) Update(entity *entity.UserEntity) (*entity.UserEntity, error) {
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

func (UserRepositoryImpl) toEntity(model *model.UserModel) (*entity.UserEntity, error) {
	id, err := uuid.FromBytes(model.Id)
	if err != nil {
		return nil, err
	}
	return &entity.UserEntity{
		Id:        id,
		Name:      model.Name,
		Password:  model.Password,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}, nil
}

func (UserRepositoryImpl) toModel(entity *entity.UserEntity) *model.UserModel {
	return &model.UserModel{
		Id:        entity.Id[:],
		Name:      entity.Name,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
