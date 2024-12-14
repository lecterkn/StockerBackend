package usecase

import (
	"fmt"
	"h11/backend/internal/stocker/domain/repository"

	"github.com/google/uuid"
)

type StoreAuthorizationUsecase struct {
	storeRepository repository.StoreRepository
}

func NewStoreAuthorizationUsecase(storeRepository repository.StoreRepository) StoreAuthorizationUsecase {
	return StoreAuthorizationUsecase{
		storeRepository,
	}
}

func (s StoreAuthorizationUsecase) IsUserRelated(storeId, userId uuid.UUID) error {
	storeEntity, err := s.storeRepository.Select(storeId)
	if err != nil {
		return err
	}
	if storeEntity.UserId == userId {
		return nil
	}
	return fmt.Errorf("user not related")
}
