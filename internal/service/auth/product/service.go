package productserviec

import (
	"context"

	"snapp-food/internal/entity"
)

type Service struct {
	repo productRepository
}

type productRepository interface {
	Create(ctx context.Context, product entity.Product) error
	GetByStoreID(ctx context.Context, storeID int) ([]entity.Product, error)
}

func New(repo productRepository) Service {
	return Service{repo: repo}
}
