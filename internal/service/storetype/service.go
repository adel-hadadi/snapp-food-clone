package storetypeservice

import (
	"context"

	"snapp-food/internal/entity"
)

type Service struct {
	repo storeTypeRepository
}

type storeTypeRepository interface {
	Get(ctx context.Context) ([]entity.StoreType, error)
}

func New(repo storeTypeRepository) Service {
	return Service{repo: repo}
}
