package service

import (
	"fmt"
	"h11/backend/internal/stocker/domain/repository"

	"github.com/google/uuid"
)

type StoreAuthorizationService struct {
	storeRepository repository.StoreRepository
}

func NewStoreAuthorizationService(storeRepository repository.StoreRepository) StoreAuthorizationService {
	return StoreAuthorizationService{
		storeRepository,
	}
}

func (s StoreAuthorizationService) IsUserRelated(storeId, userId uuid.UUID) error {
	storeEntity, err := s.storeRepository.Select(storeId)
	if err != nil {
		return err
	}
	if storeEntity.UserId == userId {
		return nil
	}
	return fmt.Errorf("user not related")
}
