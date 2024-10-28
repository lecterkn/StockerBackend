package service

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"
	"time"

	"github.com/google/uuid"
)

type ItemService struct {
	ItemRepository repository.ItemRepository
}

type ItemServiceListItem struct {
	Id uuid.UUID
	Name string
	JanCode string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ItemServiceListOutput struct {
	list []ItemServiceListItem
}


func (s ItemService) GetItems() (*ItemServiceListOutput, error) {
	entities, err := s.ItemRepository.SelectItems()
	if err != nil {
		return nil, err
	}
	return s.toOutput(entities), nil
}

func (ItemService) toOutput(entities []entity.ItemEntity) *ItemServiceListOutput {
	var list []ItemServiceListItem
	for _, entity := range entities {
		list = append(
			list, 
			ItemServiceListItem{
				Id: entity.Id,
				Name: entity.Name,
				JanCode: entity.JanCode,
				CreatedAt: entity.CreatedAt,
				UpdatedAt: entity.Updatedat,
			},
		)
	}
	return &ItemServiceListOutput{
		list: list,
	}
}