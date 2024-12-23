package productcategoryservice

import (
	"context"

	"snapp-food/internal/entity"
)

type Service struct {
	repo productCategoryRepository
}

type productCategoryRepository interface {
	Create(ctx context.Context, category entity.ProductCategory) error
	Get(ctx context.Context) ([]entity.ProductCategory, error)
}

func New(repo productCategoryRepository) Service {
	return Service{repo: repo}
}
