package storeservice

import (
	"context"

	"snapp-food/internal/entity"
)

type Service struct {
	repo              storeRepository
	storeCategoryRepo storeCategoryRepository
}

type storeRepository interface {
	Find(ctx context.Context, id int) (entity.Store, error)
	FindBySlug(ctx context.Context, slug string) (entity.Store, error)
	Create(ctx context.Context, store entity.Store) error
	Get(ctx context.Context) ([]entity.Store, error)
	ExistsByPhone(ctx context.Context, phone string) (bool, error)
}

type storeCategoryRepository interface {
	GetByStoreID(ctx context.Context, storeID int) ([]entity.StoreCategory, error)
}

func New(repo storeRepository, storeCategoryRepo storeCategoryRepository) Service {
	return Service{
		repo:              repo,
		storeCategoryRepo: storeCategoryRepo,
	}
}
