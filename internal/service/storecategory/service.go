package storecategoryservice

import (
	"context"

	"snapp-food/internal/entity"
)

type Service struct {
	repo storeCategoryRepository
}

type storeCategoryRepository interface {
	Create(ctx context.Context, category entity.StoreCategory) error
}

func New(repo storeCategoryRepository) Service {
	return Service{repo: repo}
}
