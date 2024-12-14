package service

import (
	"h11/backend/internal/stocker/domain/entity"
	"h11/backend/internal/stocker/domain/repository"
)

type ItemStockDomainService struct {
	itemRepository      repository.ItemRepository
	itemStockRepository repository.ItemStockRepository
}

func NewItemStockDomainService(itemRepository repository.ItemRepository, itemStockRepository repository.ItemStockRepository) ItemStockDomainService {
	return ItemStockDomainService{
		itemRepository,
		itemStockRepository,
	}
}

func (s ItemStockDomainService) CreateItemStock(entity *entity.ItemStockEntity) (*entity.ItemStockEntity, error) {
	itemEntity, err := s.itemRepository.Create(&entity.Item)
	if err != nil {
		return nil, err
	}
	entity.Item = *itemEntity
	itemStockEntity, err := s.itemStockRepository.Insert(entity)
	if err != nil {
		return nil, err
	}
	return itemStockEntity, nil
}
